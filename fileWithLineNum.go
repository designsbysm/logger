package timber

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const github = "github.com/designsbysm"

func fileWithLineNum(flag int) string {
	if flag == 0 {
		return ""
	}

	file, line := getCaller()
	if file == "" {
		return ""
	}

	if flag&FlagFileName != 0 {
		file = filepath.Base(file)
	} else if flag&FlagFileProject != 0 {
		current, err := os.Getwd()
		if err != nil {
			return ""
		}

		file = strings.TrimPrefix(file, current)

		if strings.Contains(file, github) {
			parts := strings.Split(file, github)
			file = parts[1]
		}
	}

	file = strings.TrimPrefix(file, "/")

	return fmt.Sprintf("%s:%d", file, line)
}

func getCaller() (string, int) {
	for i := 0; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)

		if ok && !strings.HasSuffix(file, "fileWithLineNum.go") && !strings.HasSuffix(file, "write.go") && !strings.HasSuffix(file, "helpers.go") {
			return file, line
		} else if !ok {
			return "", 0
		}
	}

	return "", 0
}
