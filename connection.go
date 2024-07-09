package main

import (
	"database/sql"
	"fmt"
	"log"
)

func (a *App) CreateConnection(database string) {

	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		"postgres", "postgres", "db", "5432", database)

	db, err := sql.Open("postgres", url)
	if err != nil {

		log.Fatal(err)
	}
	a.DB = db
}
