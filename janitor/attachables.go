package janitor

type Attachables struct {
	MinecraftAttachable `json:"minecraft:attachable"`
}

type MinecraftAttachable struct {
	Description `json:"description"`
}
