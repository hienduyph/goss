//go:build !jsoniter
// +build !jsoniter

package jsonx

import "encoding/json"

var (
	MarshalIndent = json.MarshalIndent
	NewDecoder    = json.NewDecoder
	NewEncoder    = json.NewEncoder
	Marshal       = json.Marshal
	Unmarshal     = json.Unmarshal
)
