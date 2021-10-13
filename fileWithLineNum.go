package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func fileWithLineNum(flag int) string {
	if flag == 0 {
		return ""
	}

	if _, file, line, ok := runtime.Caller(2); ok {
		if flag&FlagFileName != 0 {
			file = filepath.Base(file)
		} else {
			current, err := os.Getwd()
			if err != nil {
				return ""
			}

			file = strings.TrimPrefix(file, current)

			if strings.Contains(file, "github.com/designsbysm") {
				parts := strings.Split(file, "github.com/designsbysm")
				file = parts[1]
			}
		}

		file = strings.TrimPrefix(file, "/")

		return fmt.Sprintf("%s:%d", file, line)
	}

	return ""
}
