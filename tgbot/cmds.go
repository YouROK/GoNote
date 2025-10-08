package tgbot

import (
	"GoNote/config"

	tele "gopkg.in/telebot.v4"
)

func setupCommands(bot *tele.Bot) {
	bot.Handle("/start", startBot)
	bot.Handle("/delete", deleteNote)
}

func startBot(c tele.Context) error {
	return c.Send(config.Cfg.TGBot.StartMessage)
}

func deleteNote(c tele.Context) error {
	noteId := c.Message().Payload
	_, _, _, err := store.GetNote(noteId)
	if err != nil {
		return c.Send("Error deleting note: " + err.Error())
	}
	err = store.DeleteNote(noteId)
	if err != nil {
		return c.Send("Error deleting note: " + err.Error())
	}
	return c.Send("Note deleted")
}
