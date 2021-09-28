package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func writeToFile(title string, args ...interface{}) {
	if options.fileLevel == 0 {
		options.fileLevel = -1
	}

	if options.location == "" {
		options.location = "logs"
	}

	now := time.Now()
	year := now.Format("2006")
	month := now.Format("01")
	day := now.Format("02")

	parent := filepath.Join(options.location, year, month)
	if _, err := os.Stat(parent); os.IsNotExist(err) {
		err := os.MkdirAll(parent, 0755)
		if err != nil {
			panic(err)
		}
	}

	f, err := os.OpenFile(filepath.Join(parent, fmt.Sprintf("%s-%s-%s.log", year, month, day)), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var message []interface{}
	message = append(message, title)
	message = append(message, args...)

	fileLogger := log.New(f, "", log.LstdFlags)
	fileLogger.Println(message...)
}
