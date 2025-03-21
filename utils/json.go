package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func ReadJson(r *http.Request, data any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(data)
}

func WriteJsonError(w http.ResponseWriter, status int, message string, err any) error {
	response := map[string]any{}
	response["data"] = nil
	response["status"] = status
	response["message"] = message
	response["success"] = true
	response["error"] = err
	return WriteJson(w, status, response)
}

func WriteJsonSuccess(w http.ResponseWriter, status int, message string, data any) error {
	response := map[string]any{}
	response["data"] = data
	response["status"] = status
	response["message"] = message
	response["success"] = true
	response["error"] = nil
	return WriteJson(w, status, response)
}
