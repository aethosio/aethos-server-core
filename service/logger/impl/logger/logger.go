package logger

import (
	"errors"
	"fmt"
	"log"
	"runtime"
	"strings"
)

type Logger struct {
	Module string
}

func (l *Logger) Infof(format string, a ...interface{}) {
	log.Printf(l.format(format, a...))
}

func (l *Logger) Errorf(format string, a ...interface{}) error {
	return errors.New(l.format(format, a...))
}

func (l *Logger) Fatalf(format string, a ...interface{}) {
	log.Fatal(l.format(format, a...))
}

// NotImplemented returns an error with a message that includes the file name, line number,
// and function name of the calling function
func (l *Logger) NotImplemented() error {
	return l.Errorf("Not implemented")
}

// format formats an error which includes the file name, function, line number
// of the calling function.
func (l *Logger) format(format string, a ...interface{}) string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	fileName := frame.File
	idx := strings.Index(fileName, l.Module)
	if idx != -1 {
		fileName = fileName[idx:]
	}

	function := frame.Function
	idx = strings.Index(function, l.Module)
	if idx != -1 {
		function = function[idx:]
	}
	idx = strings.Index(function, ".")
	if idx != -1 {
		function = function[idx+1:]
	}

	s := fmt.Sprintf(format, a...)

	return fmt.Sprintf("%s %s:%d - %s", function, fileName, frame.Line, s)
}
