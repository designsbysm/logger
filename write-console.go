package logger

import (
	"log"
)

func writeToConsole(title string, args ...interface{}) {
	if options.cliLevel == 0 {
		options.cliLevel = 4
	}

	var message []interface{}
	message = append(message, title)
	message = append(message, args...)
	message = append(message, "\033[0m")

	log.Println(message...)
}
