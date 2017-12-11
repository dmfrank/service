package models

import (
	"log"

	"github.com/go-pg/pg"
)

// InitDb initializes database connection
func InitDb() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "postgres",
	})
	if db == nil {
		log.Fatal("Database connection fail")
	}
	return db
}
