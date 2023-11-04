package chunks

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Chunk(t *testing.T) {
	assert := assert.New(t)
	input := make([]string, 0, 10)
	for _, item := range strings.Split("hello world cc ee ca erac hehe", " ") {
		input = append(input, item)
	}
	out := make([][]string, 0, 10)
	for chunk := range Chunk(input, 2) {
		out = append(out, chunk)
	}
	assert.Equal(4, len(out))
	assert.Equal(2, len(out[0]))
	assert.Equal(1, len(out[len(out)-1]))
}

func Test_ChunkStream(t *testing.T) {
	assert := assert.New(t)
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
	assert.Equal(4, len(out))
	assert.Equal(2, len(out[0]))
	assert.Equal(1, len(out[len(out)-1]))
}
