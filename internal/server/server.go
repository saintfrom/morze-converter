package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/morse-converter/handlers"
)

type Server struct {
	logger *log.Logger
	server *http.Server
}

func New(logger *log.Logger) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		logger: logger,
		server: srv,
	}
}

func (s *Server) Start() error {
	s.logger.Printf("Сервер запущен на %s", s.server.Addr)
	return s.server.ListenAndServe()
}
