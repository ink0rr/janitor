package janitor

func GetKeys(description Description, keys *Set) *Set {
	for _, key := range description.Animations {
		keys.Add(key)
	}

	for _, item := range description.RenderControllers {
		switch val := item.(type) {
		case string:
			keys.Add(val)

		case map[string]interface{}:
			for key := range val {
				keys.Add(key)
			}
		}
	}
	return keys
}
