package storage

import (
	"encoding/json"
	"json-bin/bins"
	"json-bin/output"
)

type Reader interface {
	Read() ([]byte, error)
}
type Writer interface {
	Write([]byte) error
}
type Db interface {
	Reader
	Writer
}
type Storage struct {
	Bins bins.BinList
}
type StorageWithDb struct {
	Storage
	db Db
}

func New(db Db) *StorageWithDb {
	data, err := db.Read()
	if err != nil {
		return &StorageWithDb{
			Storage: Storage{
				Bins: bins.BinList{},
			},
			db: db,
		}
	}

	var storage Storage
	err = json.Unmarshal(data, &storage)
	if err != nil {
		output.PrintError("Не удалось разобрать data.json")
		return &StorageWithDb{
			Storage: Storage{
				Bins: bins.BinList{},
			},
			db: db,
		}
	}
	return &StorageWithDb{
		Storage: storage,
		db:      db,
	}
}

func (storage *StorageWithDb) Save() {
	data, err := json.Marshal(storage)
	if err != nil {
		output.PrintError("Не удалось преобразовать в JSON")
		return
	}
	err = storage.db.Write(data)
	if err != nil {
		output.PrintError(err)
	}
}
