package envx

import (
	"os"
	"strconv"
	"strings"
)

func String(key, defaultValue string) string {
	envValue := envRaw(key)
	if envValue == "" {
		return defaultValue
	}
	return envValue
}

func Int(key string, defaultValue int64) int64 {
	envValue := envRaw(key)
	if envValue == "" {
		return defaultValue
	}
	intValue, err := strconv.ParseInt(envValue, 10, 64)
	if err != nil {
		return defaultValue
	}
	return intValue
}

// Float64 gets float value
func Float(key string, defaultValue float64) float64 {
	envValue := envRaw(key)
	if envValue == "" {
		return defaultValue
	}
	v, err := strconv.ParseFloat(envValue, 64)
	if err != nil {
		return defaultValue
	}
	return v
}

// Array parses separator values into go array
func Array(key, separator string, defaultValue []string) []string {
	raw := envRaw(key)
	if raw == "" {
		return defaultValue
	}
	v := strings.Split(raw, separator)
	if len(v) == 0 {
		return defaultValue
	}
	return v
}

func envRaw(key string) string {
	return strings.TrimSpace(os.Getenv(key))
}
