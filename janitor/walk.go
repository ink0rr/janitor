package janitor

import (
	"fmt"
	"os"
	"path/filepath"
)

func Walk(path string) (paths []string) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Path not found: %s\n", path)
		return []string{}
	}
	return paths
}
