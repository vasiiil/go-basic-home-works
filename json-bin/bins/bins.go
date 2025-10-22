package bins

import (
	"math/rand/v2"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	// Name      string    `json:"name"`
}

type BinRecord struct {
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewBin(id string, private bool, name string) *Bin {
	return &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		// Name:      name,
	}
}

type BinList []Bin

func NewBinList() []Bin {
	return make([]Bin, 0)
}

func GenerateRecord() *BinRecord {
	var allowedPasswordSymbols = []rune("1234567890_-qwertyupasdfghjkzxcvbnmQWERTYYUPASDFGHJKZXCVBNM(){}[];:'\"<>,!#$&*+=")
	password := make([]rune, 32)
	for i := range 32 {
		password[i] = allowedPasswordSymbols[rand.IntN(len(allowedPasswordSymbols))]
	}
	return &BinRecord{
		Text:      string(password),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
