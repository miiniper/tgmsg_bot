package bot

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"strings"

	"go.uber.org/zap"

	"github.com/miiniper/loges"
)

type BotApi struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

const TelegramBotApi = "https://api.telegram.org/bot%s/%s?"

func NewBotApi(token, name string) (*BotApi, error) {
	return &BotApi{
		Token: token,
		Name:  name,
	}, nil
}

func (bot *BotApi) BotRequst(mod string, params url2.Values) (*http.Response, error) {
	url := fmt.Sprintf(TelegramBotApi, bot.Token, mod)
	req, err := http.NewRequest("POST", url, strings.NewReader(params.Encode()))
	if err != nil {
		loges.Loges.Error("new http request error: ", zap.Error(err))
		return &http.Response{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//	req.Header.Set("Content-Type", "application/json")
	cli := http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		loges.Loges.Error("new http resp error: ", zap.Error(err))
		return &http.Response{}, err
	}

	if resp.StatusCode != 200 {
		loges.Loges.Error("telegram api status code  not 200", zap.Int("statusCode", resp.StatusCode))
		err = errors.New("statusCode is not 200")
		return &http.Response{}, err
	}
	return resp, nil
}

//sendMessage
func (bot *BotApi) SendMsg(chatId string, text string) error {
	v := url2.Values{}
	v.Set("chat_id", chatId)
	v.Set("text", text)

	_, err := bot.BotRequst("sendMessage", v)
	if err != nil {
		loges.Loges.Error("new http resp error: ", zap.Error(err))
		return err
	}
	return nil

}

type HttpResult struct {
	Ok     bool     `json:"ok"`
	Result []Result `json:"result"`
}
type From struct {
	ID           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LanguageCode string `json:"language_code"`
}
type Chat struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	Type      string `json:"type"`
}
type Message struct {
	MessageID int    `json:"message_id"`
	From      From   `json:"from"`
	Chat      Chat   `json:"chat"`
	Date      int    `json:"date"`
	Text      string `json:"text"`
}
type Result struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

func (bot *BotApi) GetUpdates() ([]Result, error) {
	resp, err := bot.BotRequst("getUpdates", nil)
	if err != nil {
		loges.Loges.Error("new http resp error: ", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()
	var ss HttpResult
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		loges.Loges.Error("json resp body error: ", zap.Error(err))
		return nil, err
	}
	err = json.Unmarshal(body, &ss)
	if err != nil {
		loges.Loges.Error("json resp body error: ", zap.Error(err))
		return nil, err
	}
	fmt.Println(ss.Result)
	return ss.Result, nil
}

var ChatId = make(map[int]string)

//update chat
func (bot *BotApi) UpDateChatId() {
	result, _ := bot.GetUpdates()
	for _, j := range result {
		if _, ok := ChatId[j.Message.Chat.ID]; ok {
			loges.Loges.Info("chatId existed", zap.Int("chatId", j.Message.Chat.ID), zap.Any("username", j.Message.From.FirstName))
		} else {
			ChatId[j.Message.Chat.ID] = j.Message.From.FirstName
			loges.Loges.Info("update chatId successful", zap.Int("chatId", j.Message.Chat.ID), zap.String("username", j.Message.From.FirstName))
		}
	}
}

func ShowChat() {
	loges.Loges.Info("chat channel is :", zap.Any("chat", ChatId))
}

// get chat_id
func GetChatId(userName string) int {
	for i, j := range ChatId {
		if j == userName {
			return i
		}
	}
	return -1
}
