package httpx

import (
	"context"
	"net/http"
)

func NewServer() *Server {
	return &Server{Server: &http.Server{}}
}

type Server struct {
	*http.Server
}

func (s *Server) Run(ctx context.Context) error {
	errCh := make(chan error, 1)
	go func() {
		errCh <- s.ListenAndServe()
		close(errCh)
	}()

	// wait things dont
	select {
	case err, _ := <-errCh:
		return err
	case <-ctx.Done():
		return s.Shutdown(context.Background())
	}
}
