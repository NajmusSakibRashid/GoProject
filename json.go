package main

import (
	"net/http"
	"encoding/json"
	"log"
)

func respondWithError(w http.ResponseWriter, status int, message string) {
	if status > 499 {
		log.Printf("Server error: %s", message)
	}
	type errResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, status, errResponse{Error: message})
}

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(dat)
}