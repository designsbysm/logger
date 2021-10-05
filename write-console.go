package logger

import (
	"fmt"
	"time"
)

func writeToConsole(color string, args ...interface{}) {
	var message []interface{}

	format := "\033[0;90m[%s]%s"
	message = append(message, fmt.Sprintf(format, time.Now().Format("15:04:05"), color))
	message = append(message, args...)
	message = append(message, "\033[0m")

	fmt.Println(message...)
}
