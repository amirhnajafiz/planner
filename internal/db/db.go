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
)

var (
	// database config keys
	keys = map[string]string{
		"{{table}}":    tableName,
		"{{user}}":     username,
		"{{password}}": password,
		"{{host}}":     host,
		"{{port}}":     strconv.Itoa(port),
	}
)

func getConnectionKey() string {
	connectionStr := "host={{host}} port={{port}} user={{user}} password={{password}} dbname={{table}} sslmode=disable"

	for key := range keys {
		temp := os.Getenv(key)
		if temp == "" {
			temp = keys[key]
		}

		connectionStr = strings.Replace(connectionStr, key, temp, 1)
	}

	return connectionStr
}

func NewConnection() (*sql.DB, error) {
	return sql.Open(driverName, getConnectionKey())
}
