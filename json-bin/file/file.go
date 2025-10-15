package file

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

func ReadFile(fileName string) ([]byte, error) {
	err := isJson(fileName)
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(fileName)
	if err != nil {
		color.Red("Ошибка чтения")
		return nil, err
	}

	return data, nil
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		color.Red(err.Error())
		return
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		color.Red(err.Error())
		return
	}
	color.Green("Запись успешна")
}

func isJson(fileName string) error {
	ext := filepath.Ext(fileName)
	if ext != "json" {
		return errors.New("NOT_JSON")
	}
	return nil
}
