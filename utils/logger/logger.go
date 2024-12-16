package logger

import (
	"fmt"
	"log"
	"reflect"
)

func Info(message string, extra interface{}) {
	fmt.Printf("\033[34m[INFO]\033[0m %s", message)
	handleExtra(extra)
}

func Warn(message string, extra interface{}) {
	fmt.Printf("\033[33m[WARN]\033[0m %s", message)
	handleExtra(extra)
}

func Error(message string, extra interface{}) {
	fmt.Printf("\033[31m[ERROR]\033[0m %s", message)
	handleExtra(extra)
	log.Fatal("")
}

func Success(message string, extra interface{}) {
	fmt.Printf("\033[32m[SUCCESS]\033[0m %s", message)
	handleExtra(extra)
}

func handleExtra(extra interface{}) {
	if extra == nil {
		return
	}
	switch v := extra.(type) {
	case error:
		fmt.Printf(" | %v", v)
	case string:
		fmt.Printf(" | %s", v)
	default:
		fmt.Printf(" | %v (type: %s)", v, reflect.TypeOf(v))
	}
	fmt.Println()
}
