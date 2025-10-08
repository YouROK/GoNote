package tgbot

import (
	"GoNote/config"
	"GoNote/storage"
	"log"
	"time"

	tele "gopkg.in/telebot.v4"
	"gopkg.in/telebot.v4/middleware"
)

var (
	bot   *tele.Bot
	store storage.Store
)

func Start() {
	if config.Cfg.TGBot.Token == "" {
		log.Println("TGBot token is empty")
		config.Cfg.Features.DisableReportButton = true
		return
	}

	if len(config.Cfg.TGBot.AdminIds) == 0 {
		log.Println("TGBot admins is empty")
		config.Cfg.Features.DisableReportButton = true
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
		config.Cfg.Features.DisableReportButton = true
		return
	}
	log.Println("TG Bot started...")

	bot.Use(middleware.Whitelist(config.Cfg.TGBot.AdminIds...))

	setupCommands(bot)

	go bot.Start()

	SendMessageAll("Starting GoNote server...")
}

func SetStore(str storage.Store) {
	store = str
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
