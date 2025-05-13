package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/its-kos/web-crawler/utils"
	"github.com/its-kos/web-crawler/types"
)

func (s *ApiServer) getUniqueUrls(w http.ResponseWriter, r *http.Request) {

	targetUrl := r.URL.Query().Get("target")
	if targetUrl == "" {
		http.Error(w, "Missing target URL", http.StatusBadRequest)
		return
	}

	urls, err := utils.BFS(targetUrl)
	if err != nil {
		http.Error(w, "Failed to fetch URLs", http.StatusInternalServerError)
		return
	}

	log.Println(urls)

	response := types.Response{
		Domain: targetUrl,
		Pages: urls,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
