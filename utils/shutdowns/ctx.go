package shutdowns

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// NewCtx creates new Context that hooks with SIGTERM and SIGINT
func NewCtx() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sCh := make(chan os.Signal, 1)
		signal.Notify(sCh, syscall.SIGINT, syscall.SIGTERM)
		<-sCh
		cancel()
	}()
	return ctx
}
