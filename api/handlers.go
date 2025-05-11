package api

import (
	"log"
	"net/http"

	"github.com/its-kos/web-crawler/utils"
)

func (s *ApiServer) getUniqueUrls(w http.ResponseWriter, r *http.Request) {

	targetUrl := r.URL.Query().Get("target")
	if targetUrl == "" {
		http.Error(w, "Missing target URL", http.StatusBadRequest)
		return
	}

	parsedURL, err := utils.ValidateURL(targetUrl)
	if err != nil {
		http.Error(w, "Invalid target URL", http.StatusBadRequest)
		return
	}

	domain := parsedURL.Hostname() // Right now this includes www. We can remove it later if needed.

	resHTML, err := utils.FetchHTML(targetUrl)
	if err != nil {
		http.Error(w, "Failed to fetch HTML", http.StatusInternalServerError)
		return
	}

	urls, err := utils.ExtractURLs(*resHTML)
	if err != nil {
		http.Error(w, "Failed to extract URLs", http.StatusInternalServerError)
		return
	}

	log.Println(urls)

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(response)
}
