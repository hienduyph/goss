//go:build jsoniter
// +build jsoniter

package jsonx

import "github.com/json-iterator/go"

var (
	json          = jsoniter.ConfigCompatibleWithStandardLibrary
	MarshalIndent = json.MarshalIndent
	NewDecoder    = json.NewDecoder
	NewEncoder    = json.NewEncoder
	Marshal       = json.Marshal
	Unmarshal     = json.Unmarshal
)
