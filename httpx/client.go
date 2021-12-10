package httpx

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/hienduyph/goss/jsonx"
)

// Client defines cores interface that a http must impl
type Client interface {
	Do(*http.Request) (*http.Response, error)
}

// RichClient defines cores interface that a http must impl
type RichClient interface {
	Client
	Get(context.Context, string) (*http.Response, error)
	DoJSON(*http.Request, interface{}) error
	GetJSON(context.Context, string, interface{}) error
}

func NewClient(opts ...defaultClientOptFunc) RichClient {
	def := newDefaultClientOpt()
	for _, o := range opts {
		o(def)
	}
	impl := &defaultClient{
		innerClient: def.client(),
	}
	return impl
}

type defaultClient struct {
	innerClient *http.Client
}

func (c *defaultClient) Do(r *http.Request) (*http.Response, error) {
	return c.innerClient.Do(r)
}

func (c *defaultClient) Get(ctx context.Context, uri string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("contruct req failed: %w", err)
	}
	return c.Do(req)
}

func (c *defaultClient) GetJSON(ctx context.Context, path string, data interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return fmt.Errorf("contruct req failed: %w", err)
	}
	return c.DoJSON(req, data)
}

func (c *defaultClient) DoJSON(r *http.Request, data interface{}) error {
	resp, err := c.Do(r)
	if err != nil {
		return fmt.Errorf("do req failed: %w", err)
	}
	defer resp.Body.Close()
	return jsonx.NewDecoder(resp.Body).Decode(data)
}

func WithClientTimeout(dur time.Duration) defaultClientOptFunc {
	return func(dco *defaultClientOpt) {
		dco.timeout = dur
	}
}

func newDefaultClientOpt() *defaultClientOpt {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   5 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	return &defaultClientOpt{
		transport: transport,
		timeout:   30 * time.Second,
	}
}

type defaultClientOptFunc func(*defaultClientOpt)

type defaultClientOpt struct {
	transport *http.Transport
	timeout   time.Duration
}

func (d *defaultClientOpt) client() *http.Client {
	return &http.Client{Transport: d.transport, Timeout: d.timeout}
}
