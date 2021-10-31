package timber

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

func Critical(args ...interface{}) {
	Write(LevelCritical, args...)
}

func Debug(args ...interface{}) {
	Write(LevelDebug, args...)
}

func Error(args ...interface{}) {
	Write(LevelError, args...)
}

func Info(args ...interface{}) {
	Write(LevelInfo, args...)
}

func Warning(args ...interface{}) {
	Write(LevelWarning, args...)
}

func Struct(v interface{}) {
	Write(LevelDebug, fmt.Sprintf("%+v\n", v))
}

func StructToJSON(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		Error(err)
	}

	Write(LevelDebug, colorReset, string(b))
}

func StructToXML(v interface{}) {
	b, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		Error(err)
	}

	Write(LevelDebug, colorReset, string(b))
}
