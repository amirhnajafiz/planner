package db

import (
	"bytes"
	"log"
	"os/exec"

	_ "github.com/lib/pq"
)

const (
	// schema file
	psqlFilename = "schema.sql"
)

func MakeMigrate() {
	cmd := exec.Command("psql", "-U", username, "-h", databaseIP, "-d", tableName, "-a", "-f", psqlFilename)

	var out, stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error executing query. Command Output: %+v\n: %+v, %v", out.String(), stderr.String(), err)
	}
}
