package routes

import (
	"encoding/json"
	"io"
	"net/http"
	"scorpio/core"
)

func GetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(core.Bc, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}
