package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const myKey = "7374140831:AAGsgR-bpIzy-2okzrB_My_LI-JKI9Iy-EI"

var tgCommand = tgbotapi.NewSetMyCommands(
	tgbotapi.BotCommand{
		Command:     "start",
		Description: "start",
	},
	tgbotapi.BotCommand{
		Command:     "help",
		Description: "help",
	},
)

func main() {
	bot, err := tgbotapi.NewBotAPI(myKey)
	if err != nil {
		log.Panic(err)
	}
	bot.Send(tgCommand)

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			var msg tgbotapi.MessageConfig
			if update.Message.IsCommand() {
				if update.Message.Command() == "start" {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, GeneratePassword())
				} else {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Это приложение - генератор паролей")
				}
			} else {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Ты ещё кто ебанат?")
			}
			bot.Send(msg)
		}
	}
}
