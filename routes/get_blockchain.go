package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"scorpio/core"
)

func getBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(core.Bc, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

func RouteGetBlockchain(router *mux.Router) {
	router.HandleFunc("/", getBlockchain).Methods("GET")
}
