package janitor

import (
	"encoding/json"
	"fmt"
	"regexp"
)

func CheckBP() {
	keys := NewSet()

	for _, path := range Walk("BP/entities") {
		bp, _ := ReadJson[EntityBehavior](path)
		keys = GetKeys(bp.Description, keys)
		if bp.ComponentGroups == nil {
			continue
		}
		events, _ := json.Marshal(bp.Events)
		for key := range bp.ComponentGroups {
			regex := regexp.MustCompile(`"component_groups":\[.*"` + key + `".*\]`)
			if !regex.Match(events) {
				fmt.Printf("Unused component group: \"%s\" in %s\n", key, path)
				delete(bp.ComponentGroups, key)
			}
		}
	}

	CleanFiles("BP", "animations", keys)
	CleanFiles("BP", "animation_controllers", keys)
}
