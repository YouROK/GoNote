package tgbot

import (
	"GoNote/config"

	tele "gopkg.in/telebot.v4"
)

func setupCommands(bot *tele.Bot) {
	bot.Handle("/start", startBot)
}

func startBot(c tele.Context) error {
	return c.Send(config.Cfg.TGBot.StartMessage)
}
