package main

import "time"

type Bin struct {
	Id        string `json:"id"`
	Private   bool `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string `json:"name"`
}
func newBin(id string, private bool, createdAt time.Time, name string) *Bin {
	return &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: createdAt,
		Name:      name,
	}
}

type BinList = []Bin
func newBinList() []Bin {
	return make([]Bin, 0)
}

func main() {}
