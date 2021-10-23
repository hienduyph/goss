package logger

import (
	"github.com/go-logr/logr"
	"go.uber.org/zap/zapcore"
)

type Logger = logr.Logger

var Discard = logr.Discard

// logr level is invert of zap log level
// see more https://github.com/go-logr/zapr#implementation-details
const (
	LogDebugLevel int = -int(zapcore.DebugLevel)
	LogInfoLevel      = -int(zapcore.InfoLevel)
	LogWarnLevel      = -int(zapcore.WarnLevel)
	LogErrorLevel     = -int(zapcore.ErrorLevel)
	LogFatalLevel     = -int(zapcore.FatalLevel)
)
