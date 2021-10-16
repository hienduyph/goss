package shutdowns

import (
	"context"
	"errors"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_ShutdownCtx(t *testing.T) {
	assert := assert.New(t)

	timeout := time.After(3 * time.Second)
	ctx, done := NewCtx()
	defer done()
	assert.NotNil(ctx)
	go func() {
		time.Sleep(5 * time.Millisecond)
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	}()
	select {
	case <-ctx.Done():
		assert.Equal(context.Canceled, ctx.Err())
		return
	case <-timeout:
		assert.Error(errors.New("Can trap Shutdown"))
	}
}
