package janitor

func CheckRP() {
	settings := GetSettings()
	keys := NewSet()

	for _, path := range Walk("RP/attachables") {
		rp, err := ReadJson[Attachables](path)
		if err != nil {
			continue
		}
		keys = GetKeys(rp.Description, keys)
	}

	for _, path := range Walk("RP/entity") {
		rp, err := ReadJson[EntityResource](path)
		if err != nil {
			continue
		}
		keys = GetKeys(rp.Description, keys)
	}

	CleanFiles("RP", "animations", keys, settings.Remove)
	CleanFiles("RP", "animation_controllers", keys, settings.Remove)
	CleanFiles("RP", "render_controllers", keys, settings.Remove)
}
