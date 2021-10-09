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

const (
	LevelSilent = iota
	LevelCritical
	LevelError
	LevelWarning
	LevelInfo
	LevelDebug
)

type options struct {
	colorful  bool
	title     bool
	level     int
	timestamp string
	writer    io.Writer
}

var writers []options

func New(writer io.Writer, level int, colorful bool, title bool, timestamp string) {
	newWriter := options{
		colorful:  colorful,
		title:     title,
		level:     level,
		timestamp: timestamp,
		writer:    writer,
	}

	writers = append(writers, newWriter)
}
