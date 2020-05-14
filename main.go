package tgmsg_bot

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/miiniper/loges"
	"go.uber.org/zap"
)

var chatId int64 = 911000205
type Bot struct {
	Name string `json:"name"`
	BotApi *tgbotapi.BotAPI `json:"botapi"`
}

func (b Bot) SendMsg(text string) error {
	msg := tgbotapi.NewMessage(chatId, text)

	_, err := b.BotApi.Send(msg)
	if err != nil {
		loges.Loges.Error("send msg error:", zap.Error(err))
		return err
	}
	return nil

}
func NewBot(name string) Bot {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BotToken"))
	if err != nil {
		loges.Loges.Error("get token error:", zap.Error(err))
	}
	return Bot{BotApi: bot,Name: name}
}
