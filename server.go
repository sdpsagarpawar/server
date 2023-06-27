package myserver

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sdpsagarpawar/logger"
	"github.com/sdpsagarpawar/router"
)

// Server is a reusable HTTP server with routing capabilities.
type Server struct {
	Router  *router.Router
	Logger  *logger.Logger
	Timeout time.Duration
}

// NewServer creates a new instance of Server with the provided router, logger, and timeout.
func NewServer(routerInstance *router.Router, loggerInstance *logger.Logger, timeout time.Duration) *Server {
	return &Server{
		Router:  routerInstance,
		Logger:  loggerInstance,
		Timeout: timeout,
	}
}

// Start starts the HTTP server on the specified port.
func (s *Server) Start(port string) {
	s.Logger.Infof("Server started on http://localhost:%s", port)

	// Create a new HTTP server with the router
	httpServer := &http.Server{
		Addr:    ":" + port,
		Handler: s.Router,
	}

	// Start the server in a separate goroutine
	go func() {
		err := httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			s.Logger.Fatalf("Server error: %s", err)
		}
	}()

	// Call the graceful shutdown method
	s.GracefulShutdown(httpServer)
}

// GracefulShutdown shuts down the HTTP server gracefully.
func (s *Server) GracefulShutdown(server *http.Server) {
	// Create a channel to listen for an interrupt signal (e.g., SIGINT, SIGTERM)
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	// Block until an interrupt signal is received
	<-stopChan

	s.Logger.Info("Shutting down the server...")

	// Create a context with the specified timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), s.Timeout)
	defer cancel()

	// Attempt to gracefully shut down the server
	err := server.Shutdown(ctx)
	if err != nil {
		s.Logger.Errorf("Server shutdown error: %s", err)
	} else {
		s.Logger.Info("Server gracefully stopped")
	}
}
