package helpers

import (
	"os"
	"path/filepath"
)

func RelativeFilePath(filename string) (string, error) {
	if os.Getenv("GO_RUN_MODE") == "dev" {
		cwd, err := os.Getwd()
		if err != nil {
			return "", err
		}

		return filepath.Join(cwd, filename), nil
	} else {
		exePath, err := os.Executable()
		if err != nil {
			return "", err
		}

		return filepath.Join(filepath.Dir(exePath), filename), nil
	}
}
