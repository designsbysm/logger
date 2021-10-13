package logger

import (
	"errors"
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

const (
	FlagColorful = 1 << iota
	FlagFileName
	FlagFilePath
	FlagTitle
)

type options struct {
	flags     int
	level     int
	timestamp string
	writer    io.Writer
}

var writers []options

func New(writer io.Writer, level int, timestamp string, flags int) error {
	if writer == nil {
		return errors.New("io.Writer is required")
	}

	newWriter := options{
		flags:     flags,
		level:     level,
		timestamp: timestamp,
		writer:    writer,
	}

	writers = append(writers, newWriter)

	return nil
}
