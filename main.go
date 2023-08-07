package main

import (
	"ahp-be/app"
	"ahp-be/cmd"
	"ahp-be/config"
	"fmt"
)

func main() {
	mode := cmd.Init()
	fmt.Print(mode)
	cfg := config.New(mode.AppMode)

	dbMode := mode.DbMode

	api := app.New(cfg, dbMode)

	app.Run(api)
}
