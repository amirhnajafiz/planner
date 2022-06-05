package cmd

import (
	"github.com/amirhnajafiz/planner/internal/cmd/server"
	"github.com/amirhnajafiz/planner/internal/logger"
)

func Execute() {
	l := logger.New()

	server.New(l.Named("server"))
}
