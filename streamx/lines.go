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

const buf = 1024 * 1024

func ReadLines(ctx context.Context, reader io.Reader) (*Bytes, error) {
	const outSize = 10
	out := make(chan []byte, outSize)
	res := &Bytes{Stream: out}
	go func() {
		sc := bufio.NewScanner(reader)
		sc.Buffer([]byte{}, buf)
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
