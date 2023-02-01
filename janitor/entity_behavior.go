package janitor

type EntityBehavior struct {
	MinecraftEntity `json:"minecraft:entity"`
}

type MinecraftEntity struct {
	Description     `json:"description"`
	ComponentGroups map[string]interface{} `json:"component_groups"`
	Events          map[string]interface{} `json:"events"`
}
