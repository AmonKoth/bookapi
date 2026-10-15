package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func WriteJSON(writer http.ResponseWriter, status int, data any) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(data)
}

func WriteError(writer http.ResponseWriter, status int, msg string) {
	WriteJSON(writer, status, map[string]string{"error": msg})
}

func ParseID(request *http.Request) (int, error) {
	return strconv.Atoi(request.PathValue("id"))
}
