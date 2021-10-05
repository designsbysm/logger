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

func getColor(level int) string {
	color := ""

	switch level {
	case LevelError:
		color = "\033[0;31m"
	case LevelWarning:
		color = "\033[0;33m"
	case LevelInfo:
		color = "\033[0;36m"
	case LevelDebug:
		color = "\033[0;32m"
	}

	return color
}

func getTitle(level int) string {
	title := ""

	switch level {
	case LevelError:
		title = "ERROR"
	case LevelWarning:
		title = "WARN "
	case LevelInfo:
		title = "INFO "
	case LevelDebug:
		title = "DEBUG"
	}

	return title
}

func Write(level int, args ...interface{}) {
	if level <= options.fileLevel {
		title := getTitle(level)
		writeToFile(title, args...)
	}

	if level <= options.cliLevel {
		color := getColor(level)
		writeToConsole(color, args...)
	}
}
