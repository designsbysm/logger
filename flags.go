package timber

import (
	"sort"
	"strings"
)

const (
	FlagColorful = 1 << iota
	FlagFileName
	FlagFileProject
	FlagTitle
	FlagFileAbsolute
	FlagNone = 0
)

var (
	flagNames = map[int]string{
		FlagNone:         "FlagNone",
		FlagColorful:     "FlagColorful",
		FlagFileName:     "FlagFileName",
		FlagFileProject:  "FlagFileProject",
		FlagTitle:        "FlagTitle",
		FlagFileAbsolute: "FlagFileAbsolute",
	}
	flagValues = map[string]int{
		"FlagNone":         FlagNone,
		"FlagColorful":     FlagColorful,
		"FlagFileName":     FlagFileName,
		"FlagFileProject":  FlagFileProject,
		"FlagTitle":        FlagTitle,
		"FlagFileAbsolute": FlagFileAbsolute,
	}
)

func FlagsToString(flags int) string {
	found := []string{}

	for key, value := range flagNames {
		if flags&key != 0 {
			found = append(found, value)
		}
	}

	sort.Strings(found)
	return strings.Join(found, "|")
}

func StringToFlags(flags string) int {
	found := strings.Split(flags, "|")
	bitwise := 0

	for _, value := range found {
		value = strings.TrimSpace(value)
		value = strings.TrimPrefix(value, "timber.")

		bitwise = bitwise | flagValues[value]
	}

	return bitwise
}
