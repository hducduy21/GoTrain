package utils

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetURLParamNumber(w http.ResponseWriter, r *http.Request, p string) int {
	idStr := chi.URLParam(r, p)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid URL param!", http.StatusBadRequest)
		return -1
	}
	return id
}

func GetURLParamString(w http.ResponseWriter, r *http.Request, p string) string {
	idStr := chi.URLParam(r, p)
	if idStr == "" {
		http.Error(w, "Invalid URL param!", http.StatusBadRequest)
		return ""
	}
	return idStr
}
