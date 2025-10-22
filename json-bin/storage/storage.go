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
		output.PrintError(err)
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

func (storage *StorageWithDb) save() bool {
	data, err := json.Marshal(storage)
	if err != nil {
		output.PrintError("Не удалось преобразовать в JSON")
		return false
	}
	err = storage.db.Write(data)
	if err != nil {
		output.PrintError(err)
		return false
	}

	return true
}

func (storage *StorageWithDb) AddBin(bin *bins.Bin) bool {
	storage.Bins = append(storage.Bins, *bin)
	return storage.save()
}

func (storage *StorageWithDb) DeleteBin(id string) bool {
	newBins := make(bins.BinList, 0, len(storage.Bins))
	deleted := false
	for _, bin := range storage.Bins {
		if bin.Id == id {
			deleted = true
		} else {
			newBins = append(newBins, bin)
		}
	}
	if deleted {
		storage.Bins = newBins
		return storage.save()
	}
	return false
}
