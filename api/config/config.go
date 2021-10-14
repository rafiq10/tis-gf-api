package config

import "net/http"

const API_VERSION = "/api/v1"

func SetApiHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
