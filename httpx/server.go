package httpx

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hienduyph/goss/logger"
)

func NewServer() *Server {
	return &Server{Server: &http.Server{}}
}

type Server struct {
	*http.Server
	ShutdownTimeout time.Duration
}

func (s *Server) Run(ctx context.Context) error {
	errCh := make(chan error, 1)
	go func() {
		errCh <- s.ListenAndServe()
		close(errCh)
	}()

	// wait things dont
	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():

		shutdownCtx := context.Background()
		if s.ShutdownTimeout > 0 {
			var cancel context.CancelFunc
			// Give outstanding requests a deadline for completion.
			shutdownCtx, cancel = context.WithTimeout(shutdownCtx, s.ShutdownTimeout)
			defer cancel()
		}

		if e := s.Shutdown(shutdownCtx); e != nil {
			s.Close()
			return fmt.Errorf("can not shutdown graceful: %w", e)
		}
		logger.Info("Server shutdown!")
		return nil
	}
}
