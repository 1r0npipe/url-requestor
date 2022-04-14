package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func URLRequestsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "here is going to work with alert and vars: %v", vars)
}
