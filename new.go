package logger

import (
	"io"
)

var writers []loggerOptions

func New(writer io.Writer, logLevel int, colorful bool, includeTitle bool, timeFormat string) {
	options := loggerOptions{
		colorful:     colorful,
		includeTitle: includeTitle,
		logLevel:     logLevel,
		timeFormat:   timeFormat,
		writer:       writer,
	}

	writers = append(writers, options)
}
