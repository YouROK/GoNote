package tgbot

import (
	"GoNote/config"
	"log"
	"time"

	tele "gopkg.in/telebot.v4"
	"gopkg.in/telebot.v4/middleware"
)

var (
	bot *tele.Bot
)

func Start() {
	if config.Cfg.TGBot.Token == "" {
		log.Println("TGBot token is empty")
		return
	}

	if len(config.Cfg.TGBot.AdminIds) == 0 {
		log.Println("TGBot admins is empty")
		return
	}

	pref := tele.Settings{
		Token:  config.Cfg.TGBot.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	var err error

	bot, err = tele.NewBot(pref)
	if err != nil {
		log.Println("Error start tg bot:", err)
		return
	}
	log.Println("TG Bot started...")

	bot.Use(middleware.Whitelist(config.Cfg.TGBot.AdminIds...))

	setupCommands(bot)

	go bot.Start()

	SendMessageAll("Starting GoNote server...")
}

func SendMessageAll(message string) {
	for _, adminID := range config.Cfg.TGBot.AdminIds {
		recipient := &tele.User{ID: adminID}
		_, err := bot.Send(recipient, message)
		if err != nil {
			log.Printf("Failed to send message to admin %d: %v\n", adminID, err)
		}
	}
}
