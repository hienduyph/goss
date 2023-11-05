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

func GetJSON(ctx context.Context, path string, out any) error {
	return globalClient.GetJSON(ctx, path, out)
}
