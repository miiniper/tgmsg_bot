package httpd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/miiniper/tgmsg_bot/bot"

	"github.com/julienschmidt/httprouter"

	"github.com/miiniper/loges"
	"go.uber.org/zap"
)

//test chat id
var chatId string = "911000205"

// to do
type BotMsg struct {
	//	Topic string `json:"topic"`
	//	User  string `json:"user"`
	Text string `json:"text"`
	//	Token string `json:"token"`
}

func SendMsg(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	tgBot, _ := bot.NewBotApi(os.Getenv("BotToken"), "tgmsg")

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

	err = tgBot.SendMsg(chatId, bb.Text)
	if err != nil {
		loges.Loges.Error("send msg error:", zap.Error(err))
		w.Write([]byte("send error : msg error "))
		return
	}

	w.Write([]byte("ok"))
}
