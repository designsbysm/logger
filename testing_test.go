package timber

import (
	"bytes"
	"testing"
)

func TestWritersReturn(t *testing.T) {
	var buf bytes.Buffer
	writers = []options{}

	err := New(&buf, LevelAll, "", FlagNone)
	if err != nil {
		t.Fatal(err)
	}

	count := len(Writers())
	if count != 1 {
		t.Errorf("should have 1, got: %d", count)
	}
}

func TestWritersReset(t *testing.T) {
	var buf bytes.Buffer

	err := New(&buf, LevelAll, "", FlagNone)
	if err != nil {
		t.Fatal(err)
	}

	Reset()

	count := len(writers)
	if count != 0 {
		t.Errorf("should have 0, got: %d", count)
	}
}
