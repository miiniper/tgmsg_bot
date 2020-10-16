package bot

import (
	"fmt"
	"net/http"

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

func (bot *BotApi) BotRequst(url string) (*http.Response, error) {

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		loges.Loges.Error("new http request error: ", zap.Error(err))
		return &http.Response{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "_ga=GA1.2.819983868.1602751983; _gid=GA1.2.236116357.1602751983")
	cli := http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		loges.Loges.Error("new http resp error: ", zap.Error(err))
		return &http.Response{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		loges.Loges.Error("telegram api status code  not 200", zap.Int("statusCode", resp.StatusCode))
		return &http.Response{}, err
	}
	return resp, nil
}

//sendMessage
func (bot *BotApi) SendMsg(chatId string, text string) error {
	mod := fmt.Sprintf(TelegramBotApi, bot.Token, "sendMessage")
	url := mod + "chat_id=" + chatId + "&text=" + text
	_, err := bot.BotRequst(url)
	if err != nil {
		loges.Loges.Error("new http resp error: ", zap.Error(err))
		return err
	}
	return nil

}
