package log

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/thebassplayer/golang-microservices/repo-microservice/src/api/config"
)

var (
	Log *logrus.Logger
)

func init() {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		level = logrus.DebugLevel
	}
	Log = &logrus.Logger{
		Level: level,
		Out:   os.Stdout,
	}

	if config.IsProduction() {
		Log.Formatter = &logrus.JSONFormatter{}
	} else {
		Log.Formatter = &logrus.TextFormatter{
			FullTimestamp:   true,                  // Include timestamps in the logs
			ForceColors:     true,                  // Ensure color output (optional)
			TimestampFormat: "2006-01-02 15:04:05", // Custom time format
		}
	}

}

func Info(msg string, tags ...string) {
	if Log.Level < logrus.InfoLevel {
		return
	}
	Log.WithFields(parseFiels(tags...)).Info(msg)

	if Log.Level < logrus.DebugLevel {
		return
	}
	Log.WithFields(parseFiels(tags...)).Debug(msg)

	if Log.Level < logrus.ErrorLevel {
		return
	}
	Log.WithFields(parseFiels(tags...)).Error(msg)
}

func parseFiels(tags ...string) logrus.Fields {
	result := make(logrus.Fields, len(tags))
	for _, tag := range tags {
		elements := strings.Split(tag, ":")
		result[strings.TrimSpace(elements[0])] = strings.TrimSpace(elements[1])
	}
	return result
}
