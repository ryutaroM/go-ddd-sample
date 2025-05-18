package http

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func WriteError(w http.ResponseWriter, statusCode int, err error) {
	http.Error(w, err.Error(), statusCode)
}
