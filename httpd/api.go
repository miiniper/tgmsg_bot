package httpd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/miiniper/tgmsg_bot/bot"

	"github.com/julienschmidt/httprouter"

	"github.com/miiniper/loges"
	"go.uber.org/zap"
)

// to do
type BotMsg struct {
	User string `json:"user"`
	Text string `json:"text"`
	//	Token string `json:"token"`
}

var TgBot bot.BotApi

func init() {
	TgBot.Token = os.Getenv("BotToken")
	TgBot.Name = "tgmsg"
	TgBot.UpDateChatId()
}
func SendMsg(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
	chatId := bot.GetChatId(bb.User)
	chat := strconv.Itoa(chatId)
	err = TgBot.SendMsg(chat, bb.Text)
	if err != nil {
		loges.Loges.Error("send msg error:", zap.Error(err))
		w.Write([]byte("send error "))
		return
	}

	w.Write([]byte("ok"))
}

func updateChatId(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	TgBot.UpDateChatId()
	bot.ShowChat()
	w.Write([]byte("ok"))
}
func GetChatId(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	user := r.FormValue("user")
	ChatId := bot.GetChatId(user)
	var mm = map[int]string{ChatId: user}
	bb, err := json.Marshal(mm)
	if err != nil {
		loges.Loges.Error("json error:", zap.Error(err))
		w.Write([]byte("json error "))
		return
	}
	w.Write(bb)
}
