package logger

import (
	"testing"
)

func TestShouldHaveLevelUnknown(t *testing.T) {
	exemplar := levelInfo{
		color: colorWhite,
		level: LevelSilent,
		title: "UNK",
	}
	result := getLevelInfo(LevelSilent)

	if exemplar != result {
		t.Errorf("should get UKN level, got: %v", result)
	}
}

func TestShouldHaveLevelCritical(t *testing.T) {
	exemplar := levelInfo{
		color: colorCritical,
		level: LevelCritical,
		title: "CRITICAL",
	}
	result := getLevelInfo(LevelCritical)

	if exemplar != result {
		t.Errorf("should get CRITICAL level, got: %v", result)
	}
}

func TestShouldHaveLevelError(t *testing.T) {
	exemplar := levelInfo{
		color: colorError,
		level: LevelError,
		title: "ERROR",
	}
	result := getLevelInfo(LevelError)

	if exemplar != result {
		t.Errorf("should get ERROR level, got: %v", result)
	}
}

func TestShouldHaveLevelWarning(t *testing.T) {
	exemplar := levelInfo{
		color: colorWarning,
		level: LevelWarning,
		title: "WARN ",
	}
	result := getLevelInfo(LevelWarning)

	if exemplar != result {
		t.Errorf("should get WARN level, got: %v", result)
	}
}

func TestShouldHaveLevelInfo(t *testing.T) {
	exemplar := levelInfo{
		color: colorInfo,
		level: LevelInfo,
		title: "INFO ",
	}
	result := getLevelInfo(LevelInfo)

	if exemplar != result {
		t.Errorf("should get INFO level, got: %v", result)
	}
}

func TestShouldHaveLevelDebug(t *testing.T) {
	exemplar := levelInfo{
		color: colorDebug,
		level: LevelDebug,
		title: "DEBUG",
	}
	result := getLevelInfo(LevelDebug)

	if exemplar != result {
		t.Errorf("should return valid io.Writer, got: %v", result)
	}
}
