package utils

import (
	"encoding/json"
	"net/http"
)

func JsonResponseWriter(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode data to JSON", http.StatusInternalServerError)
	}
}

func JsonParse(r *http.Request, data interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		return err
	}
	return nil
}
