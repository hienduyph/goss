package envx

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetFloat(t *testing.T) {
	asr := assert.New(t)
	cases := []struct {
		val        string
		defaultVal float64
		expected   float64
	}{
		{
			val:        "1",
			defaultVal: 1.0,
			expected:   1.0,
		},
		{
			val:        "",
			defaultVal: 2.0,
			expected:   2.0,
		},
		{
			val:        "0.001",
			defaultVal: 1.0,
			expected:   0.001,
		},
		{
			val:        "xx",
			defaultVal: 2.0,
			expected:   2.0,
		},
	}
	for _, c := range cases {
		os.Setenv("Test_GetFloat", c.val)
		asr.Equal(Float("Test_GetFloat", c.defaultVal), c.expected)
	}
}

func Test_GetInt(t *testing.T) {
	asr := assert.New(t)
	cases := []struct {
		val        string
		defaultVal int64
		expected   int64
	}{
		{
			val:      "11",
			expected: 11,
		},
		{
			val:      "1000",
			expected: 1000,
		},
		{
			val:        "xx",
			defaultVal: 3,
			expected:   3,
		},
		{
			val:      "2",
			expected: 2,
		},
	}
	for _, c := range cases {
		os.Setenv("Test_GetInt", c.val)
		asr.Equal(Int("Test_GetInt", c.defaultVal), c.expected)
	}
}

func Test_String(t *testing.T) {
	asr := assert.New(t)
	cases := []struct {
		val        string
		defaultVal string
		expected   string
	}{
		{
			val:        "1",
			defaultVal: "1",
			expected:   "1",
		},
		{
			val:        "",
			defaultVal: "cc",
			expected:   "cc",
		},
		{
			val:        "1000",
			defaultVal: "1.0",
			expected:   "1000",
		},
		{
			val:        "",
			defaultVal: "",
			expected:   "",
		},
		{
			val:        "2",
			defaultVal: "2",
			expected:   "2",
		},
	}
	for _, c := range cases {
		os.Setenv("Test_String", c.val)
		asr.Equal(String("Test_String", c.defaultVal), c.expected)
	}
}
