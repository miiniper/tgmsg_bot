package httpd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/miiniper/loges"
	"go.uber.org/zap"
)

var chatId int64 = 911000205

type BotMsg struct {
	//	Topic string `json:"topic"`
	//	User  string `json:"user"`
	Text string `json:"text"`
	//	Token string `json:"token"`
}

type Bot struct {
	Name   string           `json:"name"`
	BotApi *tgbotapi.BotAPI `json:"botapi"`
}

//func (b Bot) BotSendMsg(text string) error {
//	msg := tgbotapi.NewMessage(chatId, text)
//
//	_, err := b.BotApi.Send(msg)
//	if err != nil {
//		loges.Loges.Error("send msg error:", zap.Error(err))
//		return err
//	}
//	return nil
//
//}

func NewBot(name string) (Bot, error) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BotToken"))
	if err != nil {
		loges.Loges.Error("get token error:", zap.Error(err))
		return Bot{}, err
	}
	return Bot{BotApi: bot, Name: name}, nil
}

//func InitBot() {
//	TgBot := NewBot("tgmsg")
//	fmt.Printf("bot %s created \n", TgBot.Name)
//	loges.Loges.Info("bot created ", zap.Any("botName", TgBot.Name))
//}

func SendMsg(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	tgBot, err := NewBot("tgmsg")
	if err != nil {
		w.Write([]byte("send error :ken error"))
		return
	}

	var bb BotMsg

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		loges.Loges.Error("json err :", zap.Error(err))
		w.Write([]byte("send error : json0 error "))
		return
	}

	err = json.Unmarshal(body, &bb)
	if err != nil {
		loges.Loges.Error("json err :", zap.Error(err))
		w.Write([]byte("send error : json error "))
		return
	}
	msg := tgbotapi.NewMessage(chatId, bb.Text)
	_, err = tgBot.BotApi.Send(msg)

	if err != nil {
		loges.Loges.Error("send msg error:", zap.Error(err))
		w.Write([]byte("send error : msg error "))
		return
	}

	w.Write([]byte("ok"))
}
