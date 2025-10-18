package main

import (
	"fmt"
	"json-bin/file"
	"json-bin/output"
	"json-bin/storage"
)

func main() {
	fileDb, err := file.New("data.json")
	if err != nil {
		output.PrintError(err)
		return
	}
	storage := storage.New(fileDb)
	fmt.Println(*storage)
}
