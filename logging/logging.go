package logging

import "fmt"

type logger struct{}

func (l *logger) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

type Logger interface {
	Infof(format string, args ...interface{})
}

func Info(logger Logger, format string) {
	logger.Infof(format)
}

func InitLogger() Logger {
	return &logger{}
}
