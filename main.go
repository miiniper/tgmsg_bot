package main

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/miiniper/loges"
	"go.uber.org/zap"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BotToken"))
	if err != nil {
		loges.Loges.Error("get token error:", zap.Error(err))
	}

	text := "test=test1=test2"
	msg := tgbotapi.NewMessage(911000205, text)

	for i := 0; i <= 30; i++ {
		if i > 25 {
			bot.Send(msg)
		}
	}
}
