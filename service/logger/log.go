package logger

import (
	"github.com/aethosio/aethos-server-core/service/logger/impl/logger"
)

type Logger interface {
	Infof(format string, a ...interface{})
	Errorf(format string, a ...interface{}) error
	Fatalf(format string, a ...interface{})
}

func GetLogger(module string) Logger {
	return &logger.Logger{
		Module: module,
	}
}
