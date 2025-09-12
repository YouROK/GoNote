package main

import (
	"GoNote/config"
	"GoNote/logg"
	"GoNote/web"
)

func main() {
	config.LoadConfig()
	logg.Init()
	logg.Info("Starting GoNote...")

	server := web.NewServer()
	server.Run()
}
