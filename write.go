package logger

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func Write(level int, args ...interface{}) {
	if len(writers) == 0 {
		panic(errors.New("logger.Write: no log writers defined"))
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

			message += fmt.Sprintf("%s ", time.Now().Format(w.timeFormat))
		}

		// title
		if w.colorful {
			message += levelInfo.color
		}

		if w.includeTitle {
			message += fmt.Sprintf("%s ", levelInfo.title)
		}

		// output args
		for index, item := range args {
			format := " %s"
			if index == 0 {
				format = "%s"
			}

			output := fmt.Sprint(item)
			output = strings.TrimSpace(output)

			message += fmt.Sprintf(format, output)
		}

		// clean up
		if w.colorful {
			message += colorReset
		}

		if message == "" {
			message = "nothing to log"
		}

		_, err := fmt.Fprintln(w.writer, message)
		if err != nil {
			panic(err)
		}
	}
}
