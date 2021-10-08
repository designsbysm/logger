package logger

import (
	"io"
)

const colorCritical = "\033[30;41m"
const colorDebug = "\033[0;32m"
const colorError = "\033[0;91m"
const colorInfo = "\033[0;36m"
const colorReset = "\033[0m"
const colorTime = "\033[0;90m"
const colorWarning = "\033[0;33m"
const colorWhite = "\033[0;37m"

const LevelCritical = 1
const LevelError = 2
const LevelWarning = 3
const LevelInfo = 4
const LevelDebug = 5

type options struct {
	colorful  bool
	title     bool
	logLevel  int
	timestamp string
	writer    io.Writer
}

var writers []options

func New(writer io.Writer, logLevel int, colorful bool, title bool, timestamp string) {
	newWriter := options{
		colorful:  colorful,
		title:     title,
		logLevel:  logLevel,
		timestamp: timestamp,
		writer:    writer,
	}

	writers = append(writers, newWriter)
}
