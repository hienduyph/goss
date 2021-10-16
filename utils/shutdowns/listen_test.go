package shutdowns

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testBlockListen struct {
	*httptest.Server
	closed bool
}

func (t *testBlockListen) Start() error {
	t.Server.Start()
	return nil
}

func (t *testBlockListen) Close() error {
	t.closed = true
	t.Server.Close()
	return nil
}

func Test_BlockListen(t *testing.T) {
	assert := assert.New(t)
	blockRun := &testBlockListen{
		Server: httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello, client")
		})),
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
