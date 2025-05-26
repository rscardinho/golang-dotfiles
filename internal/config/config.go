package config

import (
	"github.com/BurntSushi/toml"
	"github.com/rscardinho/golang-dotfiles/cmd/helpers"
)

type Task struct {
	Name       string `toml:"name"`
	Script     string `toml:"script"`
	Validation string `toml:"validation"`
}

type File struct {
	Packages []Task `toml:"package"`
}

func Load(filename string) (File, error) {
	filePath, err := helpers.RelativeFilePath(filename)
	if err != nil {
		return File{}, err
	}

	var file File
	if _, err := toml.DecodeFile(filePath, &file); err != nil {
		return File{}, err
	}

	return file, nil
}
