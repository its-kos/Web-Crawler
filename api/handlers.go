package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *ApiServer) getUniqueUrls(w http.ResponseWriter, r *http.Request) {

	s.requestCount++
	log.Printf("Received request for unique URLs. Total requests: %d", s.requestCount)

	// Dummy response for initial commit - testing purposes
	response := map[string]interface{}{
		"successful":    s.succesful,
		"unique_urls":   []string{"http://example.com", "http://example.org"},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}