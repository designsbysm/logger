package logger

type Options struct {
	cliLevel  int
	fileLevel int
	location  string
}

// LogError todo
const LogError = 1

// LogWarning todo
const LogWarning = 2

// LogInfo todo
const LogInfo = 3

// LogDebug todo
const LogDebug = 4

var options Options

// SetOptions todo
func SetOptions(cli int, file int, location string) {
	options = Options{
		cliLevel:  cli,
		fileLevel: file,
		location:  location,
	}
}

// Write todo
func Write(level int, args ...interface{}) {
	levelText := ""
	switch level {
	case LogError:
		levelText = "ERROR"
	case LogWarning:
		levelText = "WARN"
	case LogInfo:
		levelText = "INFO"
	case LogDebug:
		levelText = "DEBUG"
	default:
		levelText = "UNKNOWN"
	}

	if level <= options.fileLevel {
		writeToFile(levelText, args...)
	}

	if level <= options.cliLevel {
		switch level {
		case LogError:
			levelText = "\033[1;31m" + levelText
		case LogWarning:
			levelText = "\033[0;33m" + levelText
		case LogInfo:
			levelText = "\033[0;36m" + levelText
		case LogDebug:
			levelText = "\033[0;32m" + levelText
		default:
			levelText = "\033[0m" + levelText
		}

		writeToConsole(levelText, args...)
	}
}
