// provide a global opinionated logger using zap
package logger

import (
	"github.com/go-logr/zapr"
)

func CloseHook() {
	_ = global.Sync()
}

func Factory(name string) Logger {
	return zapr.NewLogger(global).WithName(name).V(LogDebugLevel)
}

var (
	logDefault   = Factory("goss")
	warnDefault  = logDefault.V(LogWarnLevel)
	debugDefault = logDefault.V(LogDebugLevel)
)

// FatalIf panic on error not empty
func FatalIf(err error, msg string, keysAndValues ...interface{}) {
	if err == nil {
		return
	}
	logDefault.Error(err, msg, keysAndValues...)
	panic(err)
}

var (
	Info  = logDefault.Info
	Error = logDefault.Error
	Debug = logDefault.Info
)
