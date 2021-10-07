package logger

import (
	"fmt"
	"strings"
	"time"
)

func Write(level int, args ...interface{}) {
	if len(writers) == 0 {
		fmt.Println("logger.Write: no log writers defined")
		return
	}

	levelInfo := getLevelInfo(level)

	for _, w := range writers {
		if level > w.logLevel {
			continue
		}

		var message string

		// timestamp
		if w.timeFormat != "" {
			if w.colorful {
				message += colorTime
			}

			message += fmt.Sprintf("[%s]", time.Now().Format(w.timeFormat))
		}

		// title
		if w.colorful {
			message += levelInfo.color
		}

		if w.includeTitle {
			message += fmt.Sprintf(" %s", levelInfo.title)
		}

		// output args
		for _, item := range args {
			output := fmt.Sprint(item)
			output = strings.TrimSpace(output)

			message += fmt.Sprintf(" %s", output)
		}

		// clean up
		if w.colorful {
			message += colorReset
		}

		if message != "" {
			fmt.Fprintln(w.writer, message)
		} else {
			fmt.Fprintln(w.writer, "logger.Write: nothing to log")
			return
		}
	}
}
