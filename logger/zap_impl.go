package logger

import (
	"log"
	"os"

	"github.com/go-logr/zapr"
	"github.com/hienduyph/goss/envx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var cfg = func() zap.Config {
	var config zap.Config
	if os.Getenv("PROD") == "yes" {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	logLevel := envx.String("LOG_LEVEL", "")
	if logLevel != "" {
		// set custom log level
		var level zapcore.Level
		if e := level.Set(logLevel); e != nil {
			panic(e)
		}
		config.Level = zap.NewAtomicLevelAt(level)
	}
	return config
}()

var global *zap.Logger = func() *zap.Logger {
	var e error
	g, e := cfg.Build()
	if e != nil {
		log.Fatalf("Unable to construct logger %s", e.Error())
	}
	return g
}()

func ZapDefault() *zap.Logger {
	return global
}

func ZapNoStack() *zap.Logger {
	c := cfg
	c.DisableStacktrace = true
	z, e := c.Build()
	if e != nil {
		panic(e)
	}
	return z
}

func ZapConfigDefault() zap.Config {
	return cfg
}

func WithZap(z *zap.Logger) Logger {
	return zapr.NewLogger(z)
}
