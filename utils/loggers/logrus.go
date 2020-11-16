package loggers

import (
	"os"

	"github.com/sirupsen/logrus"
)

func ConfigLogrus() {
	// init logger
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true

	logrus.SetFormatter(customFormatter)
	if os.Getenv("DEBUG") == "yes" {
		logrus.SetLevel(logrus.DebugLevel)
	}
}
