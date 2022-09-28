package common

import (
	"errors"
	"log"
)

var (
	ErrorRecordNotFound = errors.New("record not found")
)

func AppRecover() { // using with goroutine
	if err := recover(); err != nil {
		log.Println("Recovery error:", err)
	}
}
