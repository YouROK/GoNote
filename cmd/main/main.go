package main

import (
	"GoNote/config"
	"GoNote/tgbot"
	"GoNote/web"
	"log"
)

func main() {
	log.Println("Starting GoNote...")
	config.LoadConfig()

	tgbot.Start()
	server := web.NewServer()
	server.Run()
}
