//go:build go1.6
// +build go1.6

package shutdowns

import (
	"context"
	"os/signal"
)

// NewCtx creates new Context that hooks with SIGTERM and SIGINT
func NewCtx() (context.Context, context.CancelFunc) {
	return signal.NotifyContext(context.Background())
}
