package janitor

type Description struct {
	Animations        map[string]string `json:"animations,omitempty"`
	RenderControllers []interface{}     `json:"render_controllers,omitempty"`
}
