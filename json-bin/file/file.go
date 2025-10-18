package file

import (
	"errors"
	"json-bin/output"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

type JsonDb struct {
	fileName string
}

func New(fileName string) (*JsonDb, error) {
	err := isJson(fileName)
	if err != nil {
		return nil, err
	}
	return &JsonDb{
		fileName: fileName,
	}, nil
}

func (db *JsonDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.fileName)
	if err != nil {
		output.PrintError("Ошибка чтения:")
		output.PrintError(err)
		return []byte{}, err
	}

	return data, nil
}

func (db *JsonDb) Write(content []byte) error {
	file, err := os.Create(db.fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		return err
	}
	color.Green("Запись успешна")
	return nil
}

func isJson(fileName string) error {
	ext := filepath.Ext(fileName)
	if ext != ".json" {
		return errors.New("NOT_JSON")
	}
	return nil
}
