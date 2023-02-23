package janitor

import (
	"fmt"
	"os"
)

func CleanFiles(pack string, dir string, keys *Set, remove bool) {
	for _, path := range Walk(pack + "/" + dir) {
		dataPtr, err := ReadJson[map[string]interface{}](path)
		if err != nil {
			continue
		}
		fileData := (*dataPtr)

		data := fileData[dir].(map[string]interface{})

		prevSize := len(data)
		for key := range data {
			if !keys.Has(key) {
				fmt.Printf("Unused key: \"%s\" in %s\n", key, path)
				delete(data, key)
			}
		}
		if !remove {
			continue
		}
		currentSize := len(data)
		if currentSize == 0 {
			fmt.Printf("Empty file: %s\n", path)
			os.Remove(path)
		} else if currentSize != prevSize {
			WriteJson(path, dataPtr)
		}
	}
}
