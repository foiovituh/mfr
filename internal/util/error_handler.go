package util

import (
	"log"
)

func LogFatalIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func LogFatalIfTrue(condition bool, message string) {
	if condition {
		log.Fatal(message)
	}
}
