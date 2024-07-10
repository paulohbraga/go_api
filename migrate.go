package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/lib/pq"
)

func (a *App) Migrate() {

	m, err := migrate.New(
		"file://./files/migrations/",
		a.conn)
	if err != nil {

		log.Println(err)
	}
	if err := m.Up(); err != nil {

		log.Println(err)
	}
}
