package janitor

import (
	"encoding/json"
	"io/fs"
	"os"

	"muzzammil.xyz/jsonc"
)

func ReadJson[T interface{}](path string) (result *T, err error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	jsonc.Unmarshal([]byte(file), &result)
	return result, nil
}

func WriteJson(path string, data any) error {
	byteValue, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	os.WriteFile(path, byteValue, fs.ModeAppend)
	return nil
}
