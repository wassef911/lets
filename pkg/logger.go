// logger package provides a simple logger interface for logging messages.
package pkg

import (
	"fmt"
	"reflect"
)

type LoggerInterface interface {
	Write(message interface{})
}

var _ LoggerInterface = &Logger{}

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Write(arg interface{}) {
	switch v := arg.(type) {
	case string:
		fmt.Println("Argument is a string:", v)
	case []string:
		for _, s := range v {
			fmt.Println(s)
		}
	default:
		fmt.Println("Argument is of an unknown type:", reflect.TypeOf(arg))
	}
}
