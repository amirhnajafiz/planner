package db

import (
	"bytes"
	"database/sql"
	"log"
	"os"
	"os/exec"
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

func MakeMigrate() {
	cmd := exec.Command("psql", "-U", psqlUser, "-h", psqlHost, "-d", psqlDBName, "-a", "-f", sqlFilePath)

	var out, stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error executing query. Command Output: %+v\n: %+v, %v", out.String(), stderr.String(), err)
	}
}

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
