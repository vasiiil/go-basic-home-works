package storage

import (
	"encoding/json"
	"json-bin/bins"
	"json-bin/file"

	"github.com/fatih/color"
)

type Storage bins.BinList

func (storage Storage) Save(fileName string) {
	data, err := json.Marshal(storage)
	if err != nil {
		color.Red("Не удалось преобразовать в JSON")
		return
	}
	file.WriteFile(data, fileName)
}

func Get(fileName string) Storage {
	file, err := file.ReadFile(fileName)
	if err != nil {
		return Storage{}
	}

	var storage Storage
	err = json.Unmarshal(file, &storage)
	if err != nil {
		color.Red("Не удалось разобрать data.json")
		return Storage{}
	}
	return storage
}

