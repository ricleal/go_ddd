package infrastructure

import (
	"log"
)

type Logger struct{}

func NewLogger() *Logger {
	return new(Logger)
}

func (logger Logger) Log(args ...interface{}) {
	log.Println(args...)
}
