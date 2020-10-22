# tgmsg_bot

** This is a message gateway with telegram bot. Only text now  **

[telegramBot API](https://core.telegram.org/bots/api)

[telegramBot](https://core.telegram.org/bots)


##  API interface doc

- 1.httpcheck 

`curl -XGET http://127.0.0.1:9999/httpcheck`

- 2.send message

`curl -XPOST http://127.0.0.1:9999/sendMsg -d '{"text":"test msg yeah !","user":"miiniper"}'`

- 3 updateChatId

`curl -XGET  http://127.0.0.1:9999/updateChatId`

- 4 getChatId 

`curl -XGET http://127.0.0.1:9999/getChatId?user=miiniper`






## to do 
项目待完成
