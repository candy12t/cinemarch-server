package handler

import (
	"encoding/json"
	"net/http"
)

type HTTPError struct {
	Message string `json:"message"`
}

func NewHTTPError(message string) *HTTPError {
	return &HTTPError{
		Message: message,
	}
}

func ResponseJSON(w http.ResponseWriter, body any, status int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if body == nil {
		w.WriteHeader(status)
		return
	}

	b, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := NewHTTPError(http.StatusText(http.StatusInternalServerError))
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(status)
	w.Write(b)
}
