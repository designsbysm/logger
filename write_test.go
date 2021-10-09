package logger

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestShouldHaveMessage(t *testing.T) {
	var buf bytes.Buffer

	New(&buf, 9, false, false, "")
	Write(LevelError, "test message")

	if len(buf.Bytes()) == 0 {
		t.Errorf("should have message, got: empty")
	}
}

func TestShouldNotHaveMessage1(t *testing.T) {
	var buf bytes.Buffer

	New(&buf, 1, false, false, "")
	Write(LevelError, "test message")

	if len(buf.Bytes()) > 1 {
		t.Errorf("should not have message, got: %s", buf.String())
	}
}

func TestShouldNotHaveMessage2(t *testing.T) {
	var buf bytes.Buffer

	New(&buf, 9, false, false, "")
	Write(LevelSilent, "test message")

	if len(buf.Bytes()) > 1 {
		t.Errorf("should not have message, got: %s", buf.String())
	}
}

func TestShouldBeColorful(t *testing.T) {
	var buf bytes.Buffer

	New(&buf, 9, true, false, "[]")
	Write(LevelError, "test message")

	exemplar := []byte(fmt.Sprintf("%s[] %stest message%s\n", colorTime, colorError, colorReset))

	if !bytes.Equal(buf.Bytes(), exemplar) {
		t.Errorf("should be red, got: %s", buf.String())
	}
}

func TestShouldNotBeColorful(t *testing.T) {
	var buf bytes.Buffer

	New(&buf, 9, false, false, "")
	Write(LevelError, "test message")

	if buf.String() != "test message\n" {
		t.Errorf("should not have color, got: %s", buf.String())
	}
}

func TestShouldHaveTitle(t *testing.T) {
	var buf bytes.Buffer

	New(&buf, 9, false, true, "")
	Write(LevelError, "test message")

	if !strings.Contains(buf.String(), "ERROR") {
		t.Errorf("should have title, got: %s", buf.String())
	}
}

func TestShouldNotHaveTitle(t *testing.T) {
	var buf bytes.Buffer

	New(&buf, 9, false, false, "")
	Write(LevelError, "test message")

	if strings.Contains(buf.String(), "ERROR") {
		t.Errorf("should not have title, got: %s", buf.String())
	}
}

func TestShouldHaveTimestamp(t *testing.T) {
	var buf bytes.Buffer

	New(&buf, 9, false, false, "[15:04:05]")
	Write(LevelError, "test message")

	re := regexp.MustCompile("[0-9]{2}:[0-9]{2}:[0-9]{2}")

	if len(re.Find(buf.Bytes())) == 0 {
		t.Errorf("should have timestamp, got: %s", buf.String())
	}
}

func TestShouldNotHaveTimestamp(t *testing.T) {
	var buf bytes.Buffer

	New(&buf, 9, false, false, "")
	Write(LevelError, "test message")

	if buf.String() != "test message\n" {
		t.Errorf("should not have timestamp, got: %s", buf.String())
	}
}

func TestShouldReportNothingToLog(t *testing.T) {
	var buf bytes.Buffer

	New(&buf, 9, false, false, "")
	Write(LevelError, "")

	if buf.String() != "nothing to log\n" {
		t.Errorf("should report nothing to log, got: %s", buf.String())
	}
}

func TestShouldPanicWithoutWriters(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("did not panic")
		}
	}()

	writers = []options{}

	Write(LevelError, "test message")
}

type errorWriter struct{}

func (w errorWriter) Write(p []byte) (n int, err error) {
	return 0, errors.New("some error occurred")
}

func TestShouldPanicWriteError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("did not panic")
		}
	}()

	New(errorWriter{}, 8, false, false, "")
	Write(LevelError, "test message")
}
