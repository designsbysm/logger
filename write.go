package timber

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func Write(level int, args ...interface{}) {
	if len(writers) == 0 {
		if defaultWriter == nil {
			defaultWriter = os.Stdout
		}

		fmt.Fprintln(defaultWriter, colorError+"timber.Write Error: no log writers defined, defaulting to os.Stdout"+colorReset)

		newWriter := options{
			flags:     FlagNone,
			level:     LevelAll,
			timestamp: "[15:04:05]",
			writer:    defaultWriter,
		}

		writers = append(writers, newWriter)
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
		flagFile := w.flags&FlagFileName | w.flags&FlagFileProject | w.flags&FlagFileAbsolute
		flagTitle := w.flags&FlagTitle != 0

		// timestamp
		if w.timestamp != "" {
			if flagColorful {
				message += colorTime
			}

			message += fmt.Sprintf("%s ", time.Now().Format(w.timestamp))
		}

		// filename
		if flagFile != 0 {
			file := fileWithLineNum(flagFile)

			if file != "" {
				message += fmt.Sprintf("%s ", file)
			}
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
