package tgmsg_bot

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/miiniper/loges"
	"go.uber.org/zap"
)

var chatId int64 = 911000205

func SendMsg(text string) error {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BotToken"))
	if err != nil {
		loges.Loges.Error("get token error:", zap.Error(err))
	}

	msg := tgbotapi.NewMessage(chatId, text)

	_, err = bot.Send(msg)
	if err != nil {
		loges.Loges.Error("send msg error:", zap.Error(err))
		return err
	}
	return nil

}
