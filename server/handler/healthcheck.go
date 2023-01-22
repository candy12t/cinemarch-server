package handler

import "net/http"

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	resp := struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	}

	ResponseJSON(w, resp, http.StatusOK)
}
