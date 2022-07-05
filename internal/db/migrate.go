package db

import (
	"os"
)

const (
	PATH = "./database/todos/"

	UP   = "01_migrate_up.sql"
	DOWN = "01_migrate_down.sql"
)

// Migrate migrates our database schema
func Migrate(migrate bool) error {
	conn, err := NewConnection()
	if err != nil {
		return err
	}

	source := PATH
	if migrate {
		source = source + UP
	} else {
		source = source + DOWN
	}

	query, err := os.ReadFile(source)
	if err != nil {
		return err
	}

	_, err = conn.Exec(string(query))
	if err != nil {
		return err
	}

	return nil
}
