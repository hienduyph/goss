package shutdowns

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testBlockListen struct {
	closed    bool
	closeChan chan struct{}
}

func (t *testBlockListen) Start() error {
	select {
	case <-t.closeChan:
		return nil
	}
}

func (t *testBlockListen) Close() error {
	t.closed = true
	close(t.closeChan)
	return nil
}

func Test_BlockListen(t *testing.T) {
	assert := assert.New(t)
	blockRun := &testBlockListen{
		closeChan: make(chan struct{}, 1),
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancelFn()

	go func() {
		err := BlockListen(ctx, blockRun)
		assert.Nil(err)
	}()

	select {
	case <-time.After(100 * time.Millisecond):
		assert.Equal(true, blockRun.closed)
		return
	case <-time.After(10 * time.Second):
		assert.Nil(errors.New("Can trap Shutdown"))
	}
}
