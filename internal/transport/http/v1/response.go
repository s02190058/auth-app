package v1

import (
	"encoding/json"
	"log"
	"net/http"
)

func errorResponse(w http.ResponseWriter, code int, err error) {
	response(w, code, map[string]string{
		"error": err.Error(),
	})
}

func response(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("json.Encoder: %v", err)
	}
}
