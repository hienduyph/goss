package iox

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/hienduyph/goss/httpx"
)

var ErrUnsupportSourceProtocol = errors.New("unsupport protocol")

func Read(ctx context.Context, source string, opts ...readAllOptFn) (io.ReadCloser, error) {
	const protocolSplitter = "://"
	if !strings.Contains(source, protocolSplitter) {
		return os.Open(source)
	}
	protocol := strings.Split(source, protocolSplitter)
	opt := defaultReadAllOpt()
	for _, ofn := range opts {
		ofn(opt)
	}
	switch protocol[0] {
	case "http", "https":
		return opt.fetchHTTP(ctx, source)
	case "gs":
		// fetch google cloud storage
		// TODO: support
	case "s3":
		// fetch s3 storage
		// TODO: support
	case "hdfs":
		// fetch hdfs storage
		// TODO: support
	}
	return nil, ErrUnsupportSourceProtocol
}

// ReadAll detect the request protocol and fetch the body of this source
// currently support:
// `file://`, `` -> file system
// `http://`, `https://` -> http source
// `gs://` -> google cloud storage source
// `s3://` -> aws s3 source
// `dfs://` -> hdfs source
func ReadAll(ctx context.Context, source string, opts ...readAllOptFn) ([]byte, error) {
	r, e := Read(ctx, source, opts...)
	if e != nil {
		return nil, e
	}
	return io.ReadAll(r)
}

func defaultReadAllOpt() *readAllOpt {
	return &readAllOpt{}
}

type readAllOptFn func(*readAllOpt)

type readAllOpt struct {
}

func (r *readAllOpt) fetchHTTP(ctx context.Context, source string) (io.ReadCloser, error) {
	resp, err := httpx.Get(ctx, source)
	if err != nil {
		return nil, fmt.Errorf("http request: %w", err)
	}
	return resp.Body, nil
}
