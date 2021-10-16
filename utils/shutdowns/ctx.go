//go:build !go1.6
// +build !go1.6

package shutdowns

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// NewCtx creates new Context that hooks with SIGTERM and SIGINT
func NewCtx() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sCh := make(chan os.Signal, 1)
		signal.Notify(sCh, syscall.SIGINT, syscall.SIGTERM)
		<-sCh
	}()
	return ctx, cancel
}
