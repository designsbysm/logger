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
	} else if level < LevelCritical {
		return
	}

	levelInfo := getLevelInfo(level)

	for _, w := range writers {
		if level > w.level {
			continue
		}

		var message string
		flagColorful := w.flags&FlagColorful != 0
		flagTitle := w.flags&FlagTitle != 0

		// timestamp
		if w.timestamp != "" {
			if flagColorful {
				message += colorTime
			}

			message += fmt.Sprintf("%s ", time.Now().Format(w.timestamp))
		}

		// title
		if flagColorful {
			message += levelInfo.color
		}

		if flagTitle {
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
		if flagColorful {
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
