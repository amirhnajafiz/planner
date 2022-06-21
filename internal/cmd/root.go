package cmd

import (
	"github.com/amirhnajafiz/planner/internal/cmd/server"
	"github.com/amirhnajafiz/planner/internal/db"
	"github.com/amirhnajafiz/planner/internal/logger"
	"go.uber.org/zap"
)

func Execute() {
	l := logger.New()

	// creating a new database connection
	d, err := db.NewConnection()
	if err != nil {
		l.Fatal("database error", zap.Error(err))
	}

	server.Server{
		Db:     d,
		Logger: l.Named("server"),
	}.New()
}
