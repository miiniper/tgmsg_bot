package httpd

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/spf13/viper"

	"github.com/miiniper/tgmsg_bot/bot"

	"github.com/julienschmidt/httprouter"

	"github.com/miiniper/loges"
	"go.uber.org/zap"
)

// to do
type BotMsg struct {
	User string `json:"user"`
	Text string `json:"text"`
	File string `json:"file"`
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
		w.Write([]byte("send error : json error "))
		return
	}

	err = json.Unmarshal(body, &bb)
	if err != nil {
		loges.Loges.Error("json err :", zap.Error(err))
		w.Write([]byte("send error : json error "))
		return
	}
	chatId := bot.GetChatId(bb.User)
	if chatId < 0 {
		loges.Loges.Error("chat is not found", zap.Int("chatId", chatId))
		w.Write([]byte("chat is not found "))
		return
	}
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
	err := TgBot.UpDateChatId()
	if err != nil {
		w.Write([]byte("get chatId error"))
		return
	}
	bot.ShowChat()
	w.Write([]byte("ok"))
}
func GetChatId(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

type FileMeta struct {
	FileSha1 string `json:"file_sha_1"`
	FileName string `json:"file_name"`
	FileSize int64  `json:"file_size"`
	Location string `json:"location"`
	UploadAt string `json:"upload_at"`
}

func UploadFile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	file, head, err := r.FormFile("uploadfile")
	if err != nil {
		loges.Loges.Error("get upload file error :", zap.Error(err))
	}
	defer file.Close()
	fileMeta := FileMeta{
		FileName: head.Filename,
		Location: viper.GetString("data.tmpdir") + head.Filename,
		UploadAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	newFile, err := os.Create(fileMeta.Location)
	if err != nil {
		loges.Loges.Error("create new file error :", zap.Error(err))
		w.Write([]byte("error"))
		return
	}
	defer newFile.Close()

	fileMeta.FileSize, err = io.Copy(newFile, file)
	if err != nil {
		loges.Loges.Error("copy file error :", zap.Error(err))
		w.Write([]byte("error"))
		return
	}
	newFile.Seek(0, 0)

	w.Write([]byte("ok"))
}

func SendPhoto(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, head, err := r.FormFile("uploadfile")
	if err != nil {
		loges.Loges.Error("get upload file error :", zap.Error(err))
	}

	fileName := viper.GetString("data.tmpdir") + head.Filename
	chatId := r.FormValue("chat_id")
	err = TgBot.SendPhoto(chatId, fileName)
	if err != nil {
		loges.Loges.Error("send upload file error :", zap.Error(err))
		w.Write([]byte("err"))
		return
	}

	w.Write([]byte("ok"))
}

func SendFile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, head, err := r.FormFile("uploadfile")
	if err != nil {
		loges.Loges.Error("get upload file error :", zap.Error(err))
	}

	fileName := viper.GetString("data.tmpdir") + head.Filename
	chatId := r.FormValue("chat_id")
	err = TgBot.SendFile(chatId, fileName)
	if err != nil {
		loges.Loges.Error("send upload file error :", zap.Error(err))
		w.Write([]byte("err"))
		return
	}

	w.Write([]byte("ok"))
}
