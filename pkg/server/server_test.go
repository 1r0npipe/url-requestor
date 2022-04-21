package server

import (
	"net/http"
	"net/http/httptest"
	"time"

	"testing"

	"github.com/1r0npipe/url-requestor/pkg/config"
)

func TestRequestServer_HandleRequests(t *testing.T) {
	req, err := http.NewRequest("GET", "/request?sortKey=views&limit=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	someMock := &RequestServer{
		logger: NewLogger("debug"),
		httpServer: &http.Server{
			Addr:         "0.0.0.0:8080",
			Handler:      nil,
			ReadTimeout:  time.Duration(15) * time.Second,
			WriteTimeout: time.Duration(15) * time.Second,
		},
		config: &config.Config{
			URLs: []string{"https://raw.githubusercontent.com/assignment132/assignment/main/duckduckgo.json",
				"https://raw.githubusercontent.com/assignment132/assignment/main/google.json",
				"https://raw.githubusercontent.com/assignment132/assignment/main/wikipedia.json"},
			Workers: 3,
		},
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(someMock.HandleRequests)
	handler.ServeHTTP(rr, req)
	if http.StatusOK != rr.Code {
		t.Errorf("Server didn't respond")
	}
	limit := req.URL.Query().Get("limit")
	sortKey := req.URL.Query().Get("sortKey")
	expectedLimit := "1"
	expectedSortKey := "views"
	if limit != expectedLimit || expectedSortKey != sortKey {
		t.Errorf("wrong request parameters")
	}
}
