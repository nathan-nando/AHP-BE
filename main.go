package main

import (
	"ahp-be/app"
	"ahp-be/cmd"
	"ahp-be/config"
)

func main() {
	mode := cmd.Init()

	cfg := config.New(mode.AppMode)

	dbMode := mode.DbMode

	api := app.New(cfg, dbMode)

	app.Run(api)
}
