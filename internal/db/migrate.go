package db

import (
	"log"
	"os"
)

const (
	PATH = "./database/todos/"

	UP   = "01_migrate_up.sql"
	DOWN = "01_migrate_down.sql"
)

func Migrate(migrate bool) {
	conn, err := NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	source := PATH
	if migrate {
		source = source + UP
	} else {
		source = source + DOWN
	}

	query, err := os.ReadFile(source)
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Exec(string(query))
	if err != nil {
		log.Fatal(err)
	}
}
