package timber

import "testing"

func TestLevelToString(t *testing.T) {
	exemplar := "LevelCritical"
	result := LevelToString(LevelCritical)

	if result != exemplar {
		t.Errorf("should be %s, got: %s", exemplar, result)
	}
}

func TestStringToLevel(t *testing.T) {
	exemplar := 1
	result := StringToLevel("LevelCritical")

	if result != exemplar {
		t.Errorf("should be %d, got: %d", exemplar, result)
	}
}
