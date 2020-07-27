package envs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetOrDefault(t *testing.T) {
	assert := assert.New(t)
	key := "TEST_GOPATH"
	value := "/go/path"

	os.Setenv(key, value)
	assert.Equal(value, GetStringOrDefault(key, ""))

	assert.Equal(value, GetStringOrDefault("long-key-that-does not exist", value))
}
