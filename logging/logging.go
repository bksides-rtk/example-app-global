package logging

import (
	"fmt"
	"os"
)

var Logger logger

type logger struct {
	Out *os.File
}

func Infof(format string, args ...interface{}) {
	fmt.Fprintf(Logger.Out, format, args...)
}
