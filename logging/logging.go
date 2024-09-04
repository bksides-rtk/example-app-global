package logging

import (
	"fmt"
	"os"
)

var Logger logger

type logger struct {
	Out *os.File
}

func Info(format string) {
	fmt.Fprint(Logger.Out, format)
}
