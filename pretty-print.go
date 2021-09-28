package logger

import "encoding/json"

func PP(v interface{}) string {
	s, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err.Error()
	}

	return string(s)
}
