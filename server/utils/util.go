package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{} {
		"status"  : status,
		"message" : message,
	}
}

func Respond(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal("Error response :: ", err)
	}
}
