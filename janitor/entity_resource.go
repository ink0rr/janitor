package janitor

type EntityResource struct {
	MinecraftClientEntity `json:"minecraft:client_entity"`
}

type MinecraftClientEntity struct {
	Description `json:"description"`
}
