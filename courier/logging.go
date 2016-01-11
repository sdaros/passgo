package courier

import (
	"log"
)

type Logger interface {
	Error(v ...interface{})
	Info(v ...interface{})
}

type StandardLogger struct{}

func (st StandardLogger) Error(v ...interface{}) {
	log.Fatalln(v)
	return
}

func (st StandardLogger) Info(v ...interface{}) {
	log.Println(v)
	return
}

type NullLogger struct{}

func (nl NullLogger) Error(v ...interface{}) {}

func (nl NullLogger) Info(v ...interface{}) {}
