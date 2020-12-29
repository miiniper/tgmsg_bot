package httpd

import "time"

func Hello() {
	for {
		TgBot.SendMsg("911000205", "hello")
		time.Sleep(time.Hour)
	}

}
