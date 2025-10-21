package main

import (
	"fmt"
	"json-bin/api"
	"json-bin/config"
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
	api := api.New(config.New())
	fmt.Println(*api)
	storage := storage.New(fileDb)
	fmt.Println(*storage)
}
