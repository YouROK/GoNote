package main

import (
	"GoNote/config"
	"GoNote/localize"
	"GoNote/tgbot"
	"GoNote/web"
	"log"
)

func main() {
	log.Println("Starting GoNote...")
	config.LoadConfig()

	localize.Init()

	tgbot.Start()
	server := web.NewServer()
	server.Run()
}
