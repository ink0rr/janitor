package janitor

import (
	"encoding/json"
	"fmt"
	"regexp"
)

func CheckBP() {
	settings := GetSettings()
	keys := NewSet()

	for _, path := range Walk("BP/entities") {
		bp, err := ReadJson[EntityBehavior](path)
		if err != nil {
			continue
		}
		keys = GetKeys(bp.Description, keys)
		if bp.ComponentGroups == nil {
			continue
		}
		changed := false
		events, _ := json.Marshal(bp.Events)
		for key := range bp.ComponentGroups {
			regex := regexp.MustCompile(`"component_groups":\[.*"` + key + `".*\]`)
			if !regex.Match(events) {
				changed = true
				fmt.Printf("Unused component group: \"%s\" in %s\n", key, path)
				delete(bp.ComponentGroups, key)
			}
		}
		if settings.Remove && changed {
			WriteJson(path, bp)
		}
	}

	CleanFiles("BP", "animations", keys, settings.Remove)
	CleanFiles("BP", "animation_controllers", keys, settings.Remove)
}
