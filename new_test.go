package timber

import (
	"bytes"
	"testing"
)

func TestShouldCreateWriter(t *testing.T) {
	writers = []options{}

	var buf bytes.Buffer
	New(&buf, LevelDebug, "", FlagNone)

	if len(writers) == 0 {
		t.Errorf("should have writer, got: %v", writers)
	}
}

func TestShouldFailWithoutWriter(t *testing.T) {
	err := New(nil, LevelDebug, "", FlagNone)

	if err == nil {
		t.Errorf("should have error, got: %v", err)
	}
}
