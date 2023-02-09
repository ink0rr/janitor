package janitor

import (
	"encoding/json"
	"fmt"
	"os"
)

type Settings struct {
	Remove bool `json:"remove,omitempty"`
}

var settings *Settings

func GetSettings() *Settings {
	if settings == nil {
		args := os.Args
		if len(args) < 2 {
			args = append(args, `{}`)
		}
		err := json.Unmarshal([]byte(args[1]), &settings)
		if err != nil {
			fmt.Printf("Failed to load settings: %s\n", err)
			return &Settings{}
		}
	}
	return settings
}
