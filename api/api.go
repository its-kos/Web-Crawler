package api

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5" // Using Chi because it allows for URL parameters and easier middleware
	"github.com/go-chi/chi/v5/middleware"
)

type ApiServer struct {
	addr         string
	startTime    time.Time
}

func NewApiServer(addr string) *ApiServer {
	return &ApiServer{ addr: addr, startTime: time.Now() }
}

func (s *ApiServer) Run() error {
	router := chi.NewRouter()
	router.Use(middleware.Logger) // <--<< Logger should come before Recoverer
	router.Use(middleware.Recoverer)
	router.Use(middleware.Heartbeat("/status"))

	router.Get("/pages", s.getUniqueUrls)

	server := &http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Printf("Server is listening on %s", s.addr)

	return server.ListenAndServe()
}