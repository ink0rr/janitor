package janitor

func CheckRP() {
	keys := NewSet()

	for _, path := range Walk("RP/attachables") {
		rp, _ := ReadJson[Attachables](path)
		keys = GetKeys(rp.Description, keys)
	}

	for _, path := range Walk("RP/entity") {
		rp, _ := ReadJson[EntityResource](path)
		keys = GetKeys(rp.Description, keys)
	}

	CleanFiles("RP", "animations", keys)
	CleanFiles("RP", "animation_controllers", keys)
	CleanFiles("RP", "render_controllers", keys)
}
