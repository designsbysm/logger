package logger

type loggerOptions struct {
	cliLevel  int
	fileLevel int
	location  string
}

const LevelError = 1
const LevelWarning = 2
const LevelInfo = 3
const LevelDebug = 4

var options loggerOptions

func SetOptions(cli int, file int, location string) {
	options = loggerOptions{
		cliLevel:  cli,
		fileLevel: file,
		location:  location,
	}
}

func Write(level int, args ...interface{}) {
	levelText := ""
	switch level {
	case LevelError:
		levelText = "ERROR"
	case LevelWarning:
		levelText = "WARN "
	case LevelInfo:
		levelText = "INFO "
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
			levelText = "\033[0;31m" + levelText
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
