package main

import (
	"GoNote/config"
	"GoNote/web"
	"log"
)

func main() {
	log.Println("Starting GoNote...")
	config.LoadConfig()

	server := web.NewServer()
	server.Run()
}
