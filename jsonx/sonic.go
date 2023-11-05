package jsonx

import (
	"io"

	json "github.com/bytedance/sonic"
)

func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

// NewDecoder create a Decoder holding reader
func NewDecoder(r io.Reader) json.Decoder {
	return json.ConfigDefault.NewDecoder(r)
}

func NewEncoder(w io.Writer) json.Encoder {
	return json.ConfigDefault.NewEncoder(w)
}
