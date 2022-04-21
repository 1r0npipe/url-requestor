package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/1r0npipe/url-requestor/pkg/config"
	"github.com/1r0npipe/url-requestor/pkg/model"

	"github.com/rs/zerolog"
)

var (
	allowerParameter = map[string]bool{"sortKey": true, "limit": true}
	LIMIT_MIN        = 1
	LIMIT_MAX        = 200
	TIMEOUT_REQUEST  = 5
	LIST_DATA_SIZE   = 1024
	RETRIES          = 3
)

func pringHTTPError(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, message)
}

// checkParams the function which verifies the parameters of URL
func checkParams(w http.ResponseWriter, r *http.Request) bool {
	r.ParseForm()
	for k := range r.Form {
		if _, ok := allowerParameter[k]; !ok {
			pringHTTPError(w, "wrong request string, use only those parameters: 'sortKey' and 'limit'")
			return false
		}
	}
	return true
}

func workerHTTPRequest(w http.ResponseWriter, id int, url string, results chan<- model.BodyGenerated, logger *zerolog.Logger, wg *sync.WaitGroup) {
	defer wg.Done()
	var dataArr model.BodyGenerated
	var (
		repsErr    error
		respClient *http.Response
	)
	client := http.Client{
		Timeout: time.Duration(TIMEOUT_REQUEST) * time.Second,
	}
	logger.Debug().Msgf("Doing request to URL: %s with worker id: %d", url, id)
	for RETRIES > 0 {
		respClient, repsErr = client.Get(url)
		if repsErr != nil {
			RETRIES -= 1
			logger.Debug().Msgf("can't make request to remote side with worker %d to URL: %s", id, url)
		} else {
			break
		}
	}
	if respClient == nil {
		pringHTTPError(w, "can't make request to URL: "+url)
		logger.Debug().Msgf("can't make request to remote side")
		return
	}
	defer respClient.Body.Close()
	body, err := ioutil.ReadAll(respClient.Body)
	if err != nil {
		pringHTTPError(w, "can't read body from URL: "+url)
		logger.Debug().Msgf("can't ready body from URL")
		return
	}
	err = json.Unmarshal([]byte(body), &dataArr)
	if err != nil {
		pringHTTPError(w, "can't make json from body")
		logger.Debug().Msgf("can't make json from body")
		return
	}
	results <- dataArr

	logger.Debug().Msgf("Finished request to URL: %s with worker id: %d", url, id)

}

func URLRequestsHandler(w http.ResponseWriter, r *http.Request, logger *zerolog.Logger, conf *config.Config) {
	params := r.URL.Query()
	resultsData := make(chan model.BodyGenerated, conf.Workers)
	var wg sync.WaitGroup
	logger.Debug().Msgf("number of params: %v\n", params)

	if !checkParams(w, r) {
		logger.Debug().Msgf("issue with parameters, rejected request...")
		return
	}

	limitStr := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		pringHTTPError(w, "can't make transfer string of limit to digit")
		return
	}
	sortKey := r.URL.Query().Get("sortKey")
	if limit < LIMIT_MIN || limit > LIMIT_MAX {
		pringHTTPError(w, "wrong limit parameter: use within (1;200) interval")
		return
	}
	if sortKey != "views" && sortKey != "relevanceScore" {
		pringHTTPError(w, "wrong name of sortKey, should be one of: 'views' or 'relevanceScore'")
		return
	}
	logger.Debug().Msgf("limit: %v, sortKey: %v, from %v", limit, sortKey, r.RemoteAddr)
	for i, URL := range conf.URLs {
		wg.Add(1)
		go workerHTTPRequest(w, i, URL, resultsData, logger, &wg)
	}
	wg.Wait()
	close(resultsData)
	listOfResults := make([]model.Data, 0, LIST_DATA_SIZE)
	for arr := range resultsData {
		listOfResults = append(listOfResults, arr.Data...)
	}
	sortListResult(listOfResults, sortKey)
	if limit > len(listOfResults) {
		limit = len(listOfResults)
	}
	endResult := &model.BodyGenerated{Data: listOfResults[:limit]}
	newBody, _ := json.Marshal(endResult)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(newBody)
}

func sortListResult(array []model.Data, sortKey string) {
	switch {
	case sortKey == "relevanceScore":
		sort.Slice(array, func(i, j int) bool {
			return array[i].RelevanceScore < array[j].RelevanceScore
		})
	case sortKey == "views":
		sort.Slice(array, func(i, j int) bool {
			return float64(array[i].Views) < float64(array[j].Views)
		})
	}
}
