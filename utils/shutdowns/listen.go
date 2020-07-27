package shutdowns

import "context"

type RunnableClosable interface {
	Start() error
	Close() error
}

func BlockListen(ctx context.Context, r RunnableClosable) error {
	lisErr := make(chan error, 1)
	go func() {
		if e := r.Start(); e != nil {
			lisErr <- e
		} else {
			close(lisErr)
		}
	}()
	for {
		select {
		case err, _ := <-lisErr:
			return err
		case <-ctx.Done():
			return r.Close()
		}
	}
}
