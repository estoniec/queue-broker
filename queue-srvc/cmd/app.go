package main

import (
	"golang.org/x/net/context"
	"log/slog"
	"queue-srvc/internal/app"
	"queue-srvc/internal/config"
)

func main() {
	cfg := config.GetConfig()

	ctx := context.Background()

	a := app.NewApp(cfg)

	err := a.Run(ctx)
	if err != nil {
		slog.Error(err.Error())
		return
	}
}
