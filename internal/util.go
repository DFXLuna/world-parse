package worldparse

import (
	"os"
	"path/filepath"
)

func Exists(path string) bool {
	path, err := filepath.Abs(path)
	if err != nil {
		return false
	}
	_, err = os.Stat(path)
	return err == nil
}
