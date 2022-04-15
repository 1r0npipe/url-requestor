package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestDummyServer(t *testing.T) {
	testBody := []byte("{\"data\":[{\"url\":\"www.example.com/abc1\",\"views\":1000,\"relevanceScore\":0.1},{\"url\":\"www.example.com/abc2\",\"views\":2000,\"relevanceScore\":0.2}]}")
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write(testBody)
	}))
	defer testServer.Close()

	req, err := http.NewRequest(http.MethodGet, testServer.URL, nil)
	if err != nil {
		return
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		t.Errorf("Server didn't respond")
	}

	body, _ := ioutil.ReadAll(res.Body)

	if !reflect.DeepEqual(body, testBody) {
		t.Errorf("test failed")
	}

}
