package server

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

type RouteRegistrer func(mux *http.ServeMux)

type Server struct {
	logger *slog.Logger
	srv    *http.Server
	port   int
}

func New(logger *slog.Logger) (*Server, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	p, err := strconv.Atoi(port)
	if err != nil {
		return nil, &HTTPError{Message: "invalid port", Err: err}
	}

	server := &Server{
		logger: logger,
		srv: &http.Server{
			Addr:        ":" + port, // Default address
			Handler:     http.NewServeMux(),
			ReadTimeout: 30 * time.Second,
		},
		port: p,
	}

	return server, nil
}

func (s *Server) RegisterRoutes(handlers ...RouteRegistrer) error {
	router, ok := s.srv.Handler.(*http.ServeMux)
	if !ok {
		return &RouterError{}
	}

	for _, handler := range handlers {
		handler(router)
	}

	s.logger.Info("Routes registered")

	return nil
}

func (s *Server) Start() error {
	quit, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	s.logger.Info("Starting server", slog.Int("port", s.port))

	go func() {
		if err := s.srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			s.logger.Error("server error", slog.String("error", err.Error()))
		}

		stop() // in case server returns before signal is received
	}()

	<-quit.Done()
	s.logger.Info("Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		return &HTTPError{
			Message: "error shutting down server",
			Err:     err,
		}
	}

	s.logger.Info("Server gracefully stopped")

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("Shutting down server")
	return s.srv.Shutdown(ctx)
}
