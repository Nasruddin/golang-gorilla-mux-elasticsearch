package http_utils

import (
	"encoding/json"
	"net/http"
)

func RespondJson(w http.ResponseWriter, statusCode int, body interface{})  {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func RespondError(w http.ResponseWriter, statusCode int, err error)  {
	RespondJson(w, statusCode, err)
}

