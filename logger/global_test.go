package logger

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var errTest = errors.New("test err")

func TestFatal(t *testing.T) {
	err := errTest
	art := assert.New(t)
	art.PanicsWithError(err.Error(), func() {
		Fatal(err, "panic", "k", "v")
	})
}
