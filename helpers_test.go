package timber

import (
	"bytes"
	"fmt"
	"testing"
)

func TestHelperCritical(t *testing.T) {
	var buf bytes.Buffer

	info := getLevelInfo(LevelCritical)
	message := "test message"
	exemplar := fmt.Sprintf("%s %s\n", info.title, message)

	err := New(&buf, LevelAll, "", FlagTitle)
	if err != nil {
		t.Fatal(err)
	}

	Critical(message)

	if buf.String() != exemplar {
		t.Errorf("should be %s, got: %s", exemplar, buf.String())
	}
}

func TestHelperDebug(t *testing.T) {
	var buf bytes.Buffer

	info := getLevelInfo(LevelDebug)
	message := "test message"
	exemplar := fmt.Sprintf("%s %s\n", info.title, message)

	err := New(&buf, LevelAll, "", FlagTitle)
	if err != nil {
		t.Fatal(err)
	}

	Debug(message)

	if buf.String() != exemplar {
		t.Errorf("should be %s, got: %s", exemplar, buf.String())
	}
}

func TestHelperError(t *testing.T) {
	var buf bytes.Buffer

	info := getLevelInfo(LevelError)
	message := "test message"
	exemplar := fmt.Sprintf("%s %s\n", info.title, message)

	err := New(&buf, LevelAll, "", FlagTitle)
	if err != nil {
		t.Fatal(err)
	}

	Error(message)

	if buf.String() != exemplar {
		t.Errorf("should be %s, got: %s", exemplar, buf.String())
	}
}

func TestHelperInfo(t *testing.T) {
	var buf bytes.Buffer

	info := getLevelInfo(LevelInfo)
	message := "test message"
	exemplar := fmt.Sprintf("%s %s\n", info.title, message)

	err := New(&buf, LevelAll, "", FlagTitle)
	if err != nil {
		t.Fatal(err)
	}

	Info(message)

	if buf.String() != exemplar {
		t.Errorf("should be %s, got: %s", exemplar, buf.String())
	}
}

func TestHelperStruct(t *testing.T) {
	var buf bytes.Buffer

	info := getLevelInfo(LevelDebug)
	message := []string{}
	exemplar := fmt.Sprintf("%s %+v\n", info.title, message)

	err := New(&buf, LevelAll, "", FlagTitle)
	if err != nil {
		t.Fatal(err)
	}

	Struct(message)

	if buf.String() != exemplar {
		t.Errorf("should be %s, got: %s", exemplar, buf.String())
	}
}

func TestHelperWarning(t *testing.T) {
	var buf bytes.Buffer

	info := getLevelInfo(LevelWarning)
	message := "test message"
	exemplar := fmt.Sprintf("%s %s\n", info.title, message)

	err := New(&buf, LevelAll, "", FlagTitle)
	if err != nil {
		t.Fatal(err)
	}

	Warning(message)

	if buf.String() != exemplar {
		t.Errorf("should be %s, got: %s", exemplar, buf.String())
	}
}
