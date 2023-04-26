package httpserver

import (
	"context"
	"net/http"
	"time"
)

type Config struct {
	Port            string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
}

type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

func New(cfg *Config, handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:         ":" + cfg.Port,
			Handler:      handler,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
		},
		notify:          make(chan error, 1),
		shutdownTimeout: cfg.ShutdownTimeout,
	}
}

func (s *Server) Start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

// since we want to prevent the caller from writing to the channel,
// we use a wrapper function that converts chan error to <-chan error
func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
