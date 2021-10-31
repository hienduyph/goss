package httpx

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hienduyph/goss/logger"
)

func Run(ctx context.Context, s *http.Server, opts ...runOptFunc) error {
	opt := defaultRunOpt()
	for _, fn := range opts {
		fn(opt)
	}

	errCh := make(chan error, 1)
	go func() {
		if opt.tlsCertFile != "" && opt.tlsKeyFile != "" {
			logger.Info("Listening with TLS", "addr", s.Addr)
			errCh <- s.ListenAndServeTLS(opt.tlsCertFile, opt.tlsKeyFile)
		} else {
			logger.Info("Listening", "addr", s.Addr)
			errCh <- s.ListenAndServe()
		}
		close(errCh)
	}()

	// wait things done
	select {
	case err := <-errCh:
		return fmt.Errorf("serve failed: %w", err)

	case <-ctx.Done():
		shutdownCtx := context.Background()
		if opt.shutdownTimeout > 0 {
			var cancel context.CancelFunc
			// Give outstanding requests a deadline for completion.
			shutdownCtx, cancel = context.WithTimeout(shutdownCtx, opt.shutdownTimeout)
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

type runOptFunc func(*runopt)

func WithRunTimeout(t time.Duration) runOptFunc {
	return func(r *runopt) {
		r.shutdownTimeout = t
	}
}

func WithRunTLS(cert, key string) runOptFunc {
	return func(r *runopt) {
		r.tlsCertFile = cert
		r.tlsKeyFile = key
	}
}

func defaultRunOpt() *runopt {
	return &runopt{}
}

type runopt struct {
	shutdownTimeout time.Duration
	tlsCertFile     string
	tlsKeyFile      string
}
