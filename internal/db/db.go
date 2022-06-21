package db

import (
	"database/sql"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

const (
	// sql driver
	driverName = "postgres"

	// database variables
	tableName = "todos"
	username  = "root"
	password  = "password"

	// database connection
	host = "localhost"
	port = 5432

	// database name
	db = "planner"
)

// database config keys
func getKeys() map[string]string {
	return map[string]string{
		"{{table}}":    tableName,
		"{{user}}":     username,
		"{{password}}": password,
		"{{db}}":       db,
		"{{host}}":     host,
		"{{port}}":     strconv.Itoa(port),
	}
}

func getConnectionKey() string {
	connectionStr := "host={{host}} port={{port}} user={{user}} password={{password}} dbname={{db}} sslmode=disable"
	keys := getKeys()

	for key := range keys {
		value := os.Getenv(key)
		if value == "" {
			value = keys[key]
		}

		connectionStr = strings.Replace(connectionStr, key, value, 1)
	}

	return connectionStr
}

func NewConnection() (*sql.DB, error) {
	return sql.Open(driverName, getConnectionKey())
}
