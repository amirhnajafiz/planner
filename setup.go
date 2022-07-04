package main

import (
	"log"
	"os"

	"github.com/amirhnajafiz/planner/internal/db"
)

func main() {
	flag := false
	if os.Args[1] == "migrate" {
		flag = true
	}

	err := db.Migrate(flag)
	if err != nil {
		log.Fatal(err)
	}
}
