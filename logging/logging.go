package logging

import "fmt"

func Info(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
