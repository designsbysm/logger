package logger

func Critical(args ...interface{}) {
	Write(LevelCritical, args...)
}

func Debug(args ...interface{}) {
	Write(LevelDebug, args...)
}

func Error(args ...interface{}) {
	Write(LevelError, args...)
}

func Info(args ...interface{}) {
	Write(LevelInfo, args...)
}

func Warning(args ...interface{}) {
	Write(LevelWarning, args...)
}
