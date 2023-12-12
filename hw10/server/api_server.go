package server

import (
	"log"
	"net/http"
)

type APIServer struct {
	logger *log.Logger
	server *http.Server
	mux    *http.ServeMux
}

func NewAPIServer(logger *log.Logger) *APIServer {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	apiServer := &APIServer{
		logger: logger,
		server: server,
		mux:    mux,
	}

	return apiServer
}

func (s *APIServer) AddRoute(name string, handler http.Handler) {
	s.mux.Handle(name, handler)
}

func (s *APIServer) Start() error {
	s.logger.Println("Starting server...")

	return s.server.ListenAndServe()
}
