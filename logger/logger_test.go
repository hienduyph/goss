package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LoggerFactory(t *testing.T) {
	art := assert.New(t)
	l := Factory("Test_LoggerFactory")
	art.NotNil(l)
	l.Info("Hello world", "key", "a")
}
