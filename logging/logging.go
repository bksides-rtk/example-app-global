package logging

var Logger logger

type logger interface {
	Infof(format string, args ...interface{})
}
