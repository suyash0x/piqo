package helpers

import (
	"log"
	"os"
)

func LogError(msg string, e error) {
	if e != nil {
		log.Println(msg, e)
	}
}

func FatalOutError(msg string, e error) {
	if e != nil {
		log.Fatal(msg, e)
		os.Exit(1)
	}
}
