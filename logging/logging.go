package logging

import "fmt"

type LoggingService struct {
	logger Logger
}

var DefaultLoggingService = LoggingService{
	logger: &logger{},
}

type logger struct{}

func (l *logger) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

type Logger interface {
	Infof(format string, args ...interface{})
}

func (ls *LoggingService) Info(format string) {
	ls.logger.Infof(format)
}
