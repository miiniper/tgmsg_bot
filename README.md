# tgmsg_bot

**This is a message gateway with telegram bot. Only text now**

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

- 5 uploadFile

`use POST multipart/form-data ，body key is "uploadfile" value is file`

- 6 sendPhoto

upload photo before send photo

`use POST multipart/form-data ，body key is "uploadfile" value is file`

- 7 sendFile

upload file before send file

`use POST multipart/form-data ，body key is "uploadfile" value is file`



## 使用说明
1. 首先start bot或者发消息给bot
2. getChatId 更新用户ID
3. 可以使用发消息
4. 每次有新用户都需要更新getChatId

## to do 
项目待完成
