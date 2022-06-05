package db

import (
	"database/sql"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

const (
	// sql driver
	driverName = "postgres"

	// database variables
	tableName  = "todos"
	username   = ""
	password   = ""
	databaseIP = ""
)

var (
	// database config keys
	keys = map[string]string{
		"{{table}}":       tableName,
		"{{username}}":    username,
		"{{password}}":    password,
		"{{database_ip}}": databaseIP,
	}
)

func NewConnection() (*sql.DB, error) {
	connectionStr := "postgresql://{{username}}:{{password}}@{{database_ip}}/{{table}}?sslmode=disable"

	for key := range keys {
		temp := os.Getenv(key)
		if temp == "" {
			temp = keys[key]
		}

		connectionStr = strings.Replace(connectionStr, key, temp, 1)
	}

	db, err := sql.Open(driverName, connectionStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
