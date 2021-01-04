package server

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

var (
	healthy int32
)

// Server implements our HTTP run_server
type Server struct {
	server *http.Server
}

// NewServer creates a new HTTP Server
func NewServer(port string, h http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:           ":" + port,
			Handler:        h,
			ReadTimeout:    5 * time.Second,
			WriteTimeout:   10 * time.Second,
			IdleTimeout:    30 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

// Run starts the HTTP run_server
func (s *Server) Run() error {

	// Get hostname
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		fmt.Println("")
		log.Printf("%s - Shutdown signal received...\n", hostname)
		atomic.StoreInt32(&healthy, 0)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		s.server.SetKeepAlivesEnabled(false)
		if err := s.server.Shutdown(ctx); err != nil {
			log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	log.Printf("%s - Starting server on port %v", hostname, s.server.Addr)
	atomic.StoreInt32(&healthy, 1)
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", s.server.Addr, err)
	}

	<-done
	log.Printf("%s - Server gracefully stopped.\n", hostname)
	return nil
}
