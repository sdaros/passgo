package courier

import (
	"log"
)

type Logger interface {
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
}

type StandardLogger struct{}

func (st StandardLogger) Error(v ...interface{}) {
	log.Fatalln(v...)
	return
}

func (st StandardLogger) Errorf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
	return
}

func (st StandardLogger) Info(v ...interface{}) {
	log.Println(v...)
	return
}

func (st StandardLogger) Infof(format string, v ...interface{}) {
	log.Fatalf(format, v...)
	return
}

type NullLogger struct{}

func (nl NullLogger) Error(v ...interface{})                 {}
func (nl NullLogger) Errorf(format string, v ...interface{}) {}
func (nl NullLogger) Infof(format string, v ...interface{})  {}
func (nl NullLogger) Info(v ...interface{})                  {}
