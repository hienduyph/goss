package httpx

import (
	"context"
	"net/http"
)

// global instance
var globalClient = NewClient()

func Get(ctx context.Context, path string) (*http.Response, error) {
	return globalClient.Get(ctx, path)
}
