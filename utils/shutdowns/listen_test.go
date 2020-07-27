package shutdowns

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testBlockListen struct {
	*http.Server
	closed bool
}

func (t *testBlockListen) Start() error {
	return t.Server.ListenAndServe()
}

func (t *testBlockListen) Close() error {
	t.closed = true
	return t.Server.Close()
}

func Test_BlockListen(t *testing.T) {
	assert := assert.New(t)
	blockRun := &testBlockListen{
		Server: &http.Server{},
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
