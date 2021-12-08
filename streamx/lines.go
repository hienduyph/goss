package streamx

import (
	"bufio"
	"context"
	"io"
)

type Bytes struct {
	Stream <-chan []byte
	Err    error
}

func ReadLines(ctx context.Context, reader io.Reader) (*Bytes, error) {
	out := make(chan []byte, 10)
	res := &Bytes{Stream: out}
	go func() {
		sc := bufio.NewScanner(reader)
		sc.Buffer([]byte{}, 1024*1024)
	Main:
		for {
			select {
			case <-ctx.Done():
				// early return
				break Main
			default:
				if !sc.Scan() {
					// done
					res.Err = sc.Err()
					break Main
				}
				out <- sc.Bytes()
			}
		}
		close(out)
	}()
	return res, nil
}
