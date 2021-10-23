package timber

const (
	LevelSilent = iota
	LevelCritical
	LevelError
	LevelWarning
	LevelInfo
	LevelDebug
	LevelAll
)

var (
	levelNames = map[int]string{
		LevelSilent:   "LevelSilent",
		LevelCritical: "LevelCritical",
		LevelError:    "LevelError",
		LevelWarning:  "LevelWarning",
		LevelInfo:     "LevelInfo",
		LevelDebug:    "LevelDebug",
		LevelAll:      "LevelAll",
	}
	levelValues = map[string]int{
		"LevelSilent":   LevelSilent,
		"LevelCritical": LevelCritical,
		"LevelError":    LevelError,
		"LevelWarning":  LevelWarning,
		"LevelInfo":     LevelInfo,
		"LevelDebug":    LevelDebug,
		"LevelAll":      LevelAll,
	}
)

func LevelToString(level int) string {
	return levelNames[level]
}

func StringToLevel(level string) int {
	return levelValues[level]
}
