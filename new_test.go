package logger

import (
	"bytes"
	"testing"
)

func TestShouldCreateWriter(t *testing.T) {
	if len(writers) != 0 {
		t.Errorf("writers should empty, got: %v", writers)
	}

	var buf bytes.Buffer
	New(&buf, LevelDebug, "", 0)

	if len(writers) == 0 {
		t.Errorf("should have writer, got: %v", writers)
	}
}

func TestShouldFailWithoutWriter(t *testing.T) {
	err := New(nil, LevelDebug, "", 0)

	if err == nil {
		t.Errorf("should have error, got: %v", err)
	}
}
