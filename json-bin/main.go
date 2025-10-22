package main

import (
	"flag"
	"fmt"
	"json-bin/api"
	"json-bin/bins"
	"json-bin/config"
	"json-bin/file"
	"json-bin/output"
	"json-bin/storage"

	"github.com/fatih/color"
)

var _api *api.Api
var store *storage.StorageWithDb

func main() {
	fileDb, err := file.New("data.json")
	if err != nil {
		output.PrintError(err)
		return
	}
	_api = api.New(config.New())
	store = storage.New(fileDb)

	action, id := parseFlags()
	action(id)
}

func parseFlags() (func(string), string) {
	flagGet := flag.Bool("get", false, "Flag for call Get Bin")
	flagCreate := flag.Bool("create", false, "Flag for call Create Bin")
	flagDelete := flag.Bool("delete", false, "Flag for call Delete Bin")
	flagList := flag.Bool("list", false, "Flag for call List Bins in Storage")
	id := flag.String("id", "", "Id of Bin")
	flag.Parse()

	isGet := *flagGet
	isCreate := *flagCreate
	isDelete := *flagDelete
	isList := *flagList

	sum := 0
	idRequired := false
	var action func(string)
	if isGet {
		sum++
		action = get
		idRequired = true
	}
	if isCreate {
		sum++
		action = create
	}
	if isDelete {
		sum++
		action = delete
		idRequired = true
	}
	if isList {
		sum++
		action = list
	}

	if sum == 0 {
		panic("Empty Action")
	}
	if sum > 1 {
		panic("Too many actions")
	}
	if idRequired && *id == "" {
		panic("For Get and Delete methods id is required")
	}

	return action, *id
}

func get(id string) {
	err := _api.Get(id)
	if err != nil {
		output.PrintError("Error in get")
		output.PrintError(err)
	}
}

func create(_ string) {
	binRecord := bins.GenerateRecord()
	bin, err := _api.Create(binRecord)
	if err != nil {
		output.PrintError("Error in create")
		output.PrintError(err)
		return
	}

	color.Cyan("Created bin:\n")
	output.PrintJson(bin)

	if bin != nil {
		added := store.AddBin(bin)
		if added {
			color.Green("Добавлено")
		} else {
			color.Yellow("Не добавлено")
		}
	}
}

func delete(id string) {
	err := _api.Delete(id)
	if err != nil {
		output.PrintError("Error in delete")
		color.Yellow("Не удалено из апи")
		output.PrintError(err)
		return
	}

	deleted := store.DeleteBin(id)
	if deleted {
		color.Green("Удалено")
	} else {
		color.Yellow("Не удалено из стора")
	}

}

func list(_ string) {
	fmt.Println()
	color.Cyan("Storage bins:\n")
	output.PrintJson(store.Bins)
}
