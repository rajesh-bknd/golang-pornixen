package database

import (
	"fmt"
	"github.com/go-pg/pg"
)

var DBInstance *pg.DB

func init() {
	DBInstance = pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "postgres",
		Database: "content_info",
	})
	if DBInstance == nil {
		panic(`failed to connect to database`)
	}
	fmt.Println(`Connected...`)
}
