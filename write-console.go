package logger

import (
	"fmt"
	"time"
)

func writeToConsole(title string, args ...interface{}) {
	if options.cliLevel == 0 {
		options.cliLevel = 4
	}

	var message []interface{}
	message = append(message, fmt.Sprintf("\033[0;90m[%s]", time.Now().Format("15:04:05")))
	message = append(message, title)
	message = append(message, args...)
	message = append(message, "\033[0m")

	fmt.Println(message...)
}
