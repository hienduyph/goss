package jsonx

import (
	json "github.com/bytedance/sonic"
)

func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

var NewDecoder = json.ConfigDefault.NewDecoder
