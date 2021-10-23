package timber

import (
	"strings"
	"testing"
)

func TestGetCaller(t *testing.T) {
	file, line := getCaller()
	exemplarFile := "fileWithLineNum_test.go"
	exemplarLine := 9

	if !strings.HasSuffix(file, exemplarFile) {
		t.Errorf("should end with %s, got: %s", exemplarFile, file)
	}

	if line != exemplarLine {
		t.Errorf("should be %d, got: %d", exemplarLine, line)
	}
}

func TestFileAbsoluteFlag(t *testing.T) {
	exemplar := "timber/fileWithLineNum_test.go"
	result := fileWithLineNum(FlagFileAbsolute)

	if !strings.Contains(result, exemplar) {
		t.Errorf("should include %s, got: %s", exemplar, result)
	}
}

func TestFileNameFlag(t *testing.T) {
	exemplar := "fileWithLineNum_test.go"
	result := fileWithLineNum(FlagFileName)

	if !strings.HasPrefix(result, exemplar) {
		t.Errorf("should start with %s, got: %s", exemplar, result)
	}
}

func TestFileProjectFlag(t *testing.T) {
	exemplar := "fileWithLineNum_test.go"
	result := fileWithLineNum(FlagFileProject)

	if !strings.HasPrefix(result, exemplar) {
		t.Errorf("should start with %s, got: %s", exemplar, result)
	}
}

func TestMissingFlags(t *testing.T) {
	result := fileWithLineNum(0)

	if result != "" {
		t.Errorf("should be empty, got: %s", result)
	}
}
