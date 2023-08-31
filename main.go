package main

import (
	"log/slog"
	"os"

	"github.com/dimitrilw/go-todo-app/database"
	"github.com/dimitrilw/go-todo-app/router"
)

func main() {
	logLevel := &slog.LevelVar{}  // default = Info
	logLevel.Set(slog.LevelDebug) // uncomment this line for Debug messages

	slog.SetDefault(slog.New(slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: logLevel,
		},
	)))
	slog.Debug("logger initialized")

	database.InitDB()
	database.AddTestData()

	r := router.SetupRouter()
	r.Run()
}
