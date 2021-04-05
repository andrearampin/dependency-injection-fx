package clients

import "math/rand"

type DB interface {
	Get() string
}

var names = []string{
	"Andrea",
	"Giancarlo",
	"Ray",
	"Scott",
	"Tao",
	"Antonio",
	"Sumeet",
	"Kevin",
	"Kabir",
	"Jie Ming",
}

type implementDB struct{}

func NewDB() DB {
	return &implementDB{}
}

func (i *implementDB) Get() string {
	return names[rand.Intn(len(names))]
}
