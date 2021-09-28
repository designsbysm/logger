package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Options todo
type Options struct {
	cliLevel  int
	fileLevel int
}

// LogError todo
const LogError = 1

// LogWarning todo
const LogWarning = 2

// LogInfo todo
const LogInfo = 3

// LogDebug todo
const LogDebug = 4

var options Options

// SetOptions todo
func SetOptions(cli int, file int) {
	options = Options{
		cliLevel:  cli,
		fileLevel: file,
	}
}

// package main

// import "fmt"

// func adder() func(int) int {
// 	sum := 0
// 	return func(x int) int {
// 		sum += x
// 		return sum
// 	}
// }

// func main() {
// 	pos, neg := adder(), adder()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(
// 			pos(i),
// 			neg(-2*i),
// 		)
// 	}
// }

// 	if cliLevel == 0 {
// 	cliLevel = 4
// }

// if fileLevel == 0 {
// 	fileLevel = 3
// }

// func SetTest(new int) {
// 	test = new
// }

// func GetTest() int {
// 	return test
// }

// New todo
// func New(cliLevel int, fileLevel int) func(int, ...interface{}) {

// 	if cliLevel == 0 {
// 		cliLevel = 4
// 	}

// 	// fileLevel, _ := strconv.Atoi(os.Getenv("LOG_FILE_LEVEL"))
// 	if fileLevel == 0 {
// 		fileLevel = 3
// 	}

// 	return func(level int, args ...interface{}) {
// 		// TODO: add log level configurations

// 		// cliLevel, _ := strconv.Atoi(os.Getenv("LOG_CLI_LEVEL"))

// 		levelText := ""
// 		switch level {
// 		case LogError:
// 			levelText = "ERROR"
// 		case LogWarning:
// 			levelText = "WARN"
// 		case LogInfo:
// 			levelText = "INFO"
// 		case LogDebug:
// 			levelText = "DEBUG"
// 		default:
// 			levelText = "UNKNOWN"
// 		}

// 		if level <= fileLevel {
// 			now := time.Now()
// 			year := now.Format("2006")
// 			month := now.Format("01")
// 			day := now.Format("02")

// 			// TODO: add log location configuration
// 			parent := filepath.Join("logs", year, month)
// 			if _, err := os.Stat(parent); os.IsNotExist(err) {
// 				err := os.MkdirAll(parent, 0755)
// 				if err != nil {
// 					panic(err)
// 				}
// 			}

// 			f, err := os.OpenFile(filepath.Join(parent, fmt.Sprintf("%s-%s-%s.log", year, month, day)), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 			if err != nil {
// 				panic(err)
// 			}
// 			defer f.Close()

// 			var message []interface{}
// 			message = append(message, levelText)
// 			message = append(message, args...)

// 			fileLogger := log.New(f, "", log.LstdFlags)
// 			fileLogger.Println(message...)
// 		}

// 		if level <= cliLevel {
// 			switch level {
// 			case LogError:
// 				levelText = "\033[1;31m" + levelText
// 			case LogWarning:
// 				levelText = "\033[0;33m" + levelText
// 			case LogInfo:
// 				levelText = "\033[0;36m" + levelText
// 			case LogDebug:
// 				levelText = "\033[0;32m" + levelText
// 			default:
// 				levelText = "\033[0m" + levelText
// 			}

// 			var message []interface{}
// 			message = append(message, levelText)
// 			message = append(message, args...)
// 			message = append(message, "\033[0m")

// 			log.Println(message...)
// 		}
// 	}
// }

// Write todo
func Write(level int, args ...interface{}) {
	if options.cliLevel == 0 {
		options.cliLevel = 4
	}

	if options.fileLevel == 0 {
		options.fileLevel = -1
	}

	levelText := ""
	switch level {
	case LogError:
		levelText = "ERROR"
	case LogWarning:
		levelText = "WARN"
	case LogInfo:
		levelText = "INFO"
	case LogDebug:
		levelText = "DEBUG"
	default:
		levelText = "UNKNOWN"
	}

	if level <= options.fileLevel {
		now := time.Now()
		year := now.Format("2006")
		month := now.Format("01")
		day := now.Format("02")

		// TODO: add log location configuration
		parent := filepath.Join("logs", year, month)
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
		message = append(message, levelText)
		message = append(message, args...)

		fileLogger := log.New(f, "", log.LstdFlags)
		fileLogger.Println(message...)
	}

	if level <= options.cliLevel {
		switch level {
		case LogError:
			levelText = "\033[1;31m" + levelText
		case LogWarning:
			levelText = "\033[0;33m" + levelText
		case LogInfo:
			levelText = "\033[0;36m" + levelText
		case LogDebug:
			levelText = "\033[0;32m" + levelText
		default:
			levelText = "\033[0m" + levelText
		}

		var message []interface{}
		message = append(message, levelText)
		message = append(message, args...)
		message = append(message, "\033[0m")

		log.Println(message...)
	}
}
