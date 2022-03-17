package timber

import (
	"errors"
	"io"
)

type options struct {
	flags     int
	level     int
	timestamp string
	writer    io.Writer
}

var defaultWriter io.Writer
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
