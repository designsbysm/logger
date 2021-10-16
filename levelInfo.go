package timber

type levelInfo struct {
	color string
	level int
	title string
}

func getLevelInfo(level int) levelInfo {
	info := levelInfo{
		color: colorWhite,
		level: level,
		title: "UNK",
	}

	switch level {
	case LevelCritical:
		info.color = colorCritical
		info.title = "CRITICAL"
	case LevelError:
		info.color = colorError
		info.title = "ERROR"
	case LevelWarning:
		info.color = colorWarning
		info.title = "WARN "
	case LevelInfo:
		info.color = colorInfo
		info.title = "INFO "
	case LevelDebug:
		info.color = colorDebug
		info.title = "DEBUG"
	}

	return info
}
