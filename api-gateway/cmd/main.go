package main

import (
	"api-gateway/internal/app"
	"api-gateway/internal/config"
	"context"
)

func main() {
	context := context.Background()
	c := config.GetConfig()

	app := app.NewApp(c)
	app.Run(context)
}
