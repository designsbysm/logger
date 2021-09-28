package logger

type loggerOptions struct {
	cliLevel  int
	fileLevel int
	location  string
}

// LevelError todo
const LevelError = 1

// LevelWarning todo
const LevelWarning = 2

// LevelInfo todo
const LevelInfo = 3

// LevelDebug todo
const LevelDebug = 4

var options loggerOptions

// SetOptions todo
func SetOptions(cli int, file int, location string) {
	options = loggerOptions{
		cliLevel:  cli,
		fileLevel: file,
		location:  location,
	}
}

// Write todo
func Write(level int, args ...interface{}) {
	levelText := ""
	switch level {
	case LevelError:
		levelText = "ERROR"
	case LevelWarning:
		levelText = "WARN"
	case LevelInfo:
		levelText = "INFO"
	case LevelDebug:
		levelText = "DEBUG"
	default:
		levelText = "UNKNOWN"
	}

	if level <= options.fileLevel {
		writeToFile(levelText, args...)
	}

	if level <= options.cliLevel {
		switch level {
		case LevelError:
			levelText = "\033[1;31m" + levelText
		case LevelWarning:
			levelText = "\033[0;33m" + levelText
		case LevelInfo:
			levelText = "\033[0;36m" + levelText
		case LevelDebug:
			levelText = "\033[0;32m" + levelText
		default:
			levelText = "\033[0m" + levelText
		}

		writeToConsole(levelText, args...)
	}
}
