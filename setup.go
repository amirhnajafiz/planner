package main

import (
	"log"

	"github.com/amirhnajafiz/planner/internal/db"
)

func main() {
	err := db.Migrate(true)
	if err != nil {
		log.Fatal(err)
	}
}
