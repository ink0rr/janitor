package janitor

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"

	"muzzammil.xyz/jsonc"
)

func ReadJson[T interface{}](path string) (result *T, err error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = jsonc.Unmarshal([]byte(file), &result)
	if err != nil {
		fmt.Printf("Failed to parse %s: %s\n", path, err)
		return nil, err
	}
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
