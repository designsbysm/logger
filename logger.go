package logger

import (
	"io"
)

const colorDebug = "\033[0;32m"
const colorError = "\033[0;31m"
const colorInfo = "\033[0;36m"
const colorReset = "\033[0m"
const colorTime = "\033[0;90m"
const colorWarning = "\033[0;33m"
const colorWhite = "\033[0;37m"

const LevelError = 1
const LevelWarning = 2
const LevelInfo = 3
const LevelDebug = 4

type loggerOptions struct {
	colorful     bool
	includeTitle bool
	logLevel     int
	timeFormat   string
	writer       io.Writer
}
