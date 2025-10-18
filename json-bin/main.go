package main

import (
	"fmt"
	"json-bin/storage"
)

func main() {
	storage := storage.Get("data.json")
	fmt.Println(storage)
}
