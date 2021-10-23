package timber

import "testing"

func TestFlagsToString(t *testing.T) {
	exemplar := "FlagColorful|FlagFileAbsolute"
	result := FlagsToString(FlagFileAbsolute | FlagColorful)

	if result != exemplar {
		t.Errorf("should be %s, got: %s", exemplar, result)
	}
}

func TestStringToFlags(t *testing.T) {
	exemplar := FlagFileAbsolute | FlagColorful
	result := StringToFlags("FlagColorful|FlagFileAbsolute")

	if result != exemplar {
		t.Errorf("should be %d, got: %d", exemplar, result)
	}
}
