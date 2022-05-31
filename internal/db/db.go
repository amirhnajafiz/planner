package db

import (
	"database/sql"
	"strings"
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
		connectionStr = strings.Replace(connectionStr, key, keys[key], 1)
	}

	db, err := sql.Open(driverName, connectionStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
