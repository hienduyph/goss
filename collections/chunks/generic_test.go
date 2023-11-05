package chunks

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Chunk(t *testing.T) {
	art := assert.New(t)
	input := make([]string, 0, 10)
	input = append(input, strings.Split("hello world cc ee ca erac hehe", " ")...)
	out := make([][]string, 0, 10)
	for chunk := range Chunk(input, 2) {
		out = append(out, chunk)
	}
	art.Equal(4, len(out))
	art.Equal(2, len(out[0]))
	art.Equal(1, len(out[len(out)-1]))
}

func Test_ChunkStream(t *testing.T) {
	art := assert.New(t)
	input := make(chan string, 10)
	go func() {
		for _, w := range strings.Split("hello world cc ee ca erac hehe", " ") {
			input <- w
		}
		close(input)
	}()
	out := make([][]string, 0, 10)
	for chunk := range ChunkStream(input, 2) {
		out = append(out, chunk)
	}
	art.Equal(4, len(out))
	art.Equal(2, len(out[0]))
	art.Equal(1, len(out[len(out)-1]))
}
