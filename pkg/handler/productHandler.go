package handler

import (
	"encoding/json"
	"net/http"
)

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	resp, _ := json.Marshal("Hello World")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
