package main

import (
	"github.com/beego/beego/v2/server/web"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"teaseBot/controllers"
)

const telegramTokenKey = "TG_TOKEN"
const chadGptTokenKey = "CG_TOKEN"
const triggerMessage = "!ботик давай"

func main() {
	chadGptToken := os.Getenv(chadGptTokenKey)
	logger := log.New(os.Stdout, "INFO: ", 0)
	chadGpt := NewChadGpt(chadGptToken, logger)
	token := os.Getenv(telegramTokenKey)
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	lastMessage := ""
	lastMessageId := 0

	go func() {
		web.Router("/", &controllers.MainController{})
		web.Run(":8080")
	}()

	for update := range updates {
		if update.Message != nil {
			if update.Message.Text == triggerMessage {
				messageReply := "Брат, чет отвалилось, не смогу получить ответ"
				resp, err := chadGpt.GetAnswer(lastMessage)
				if nil == err {
					messageReply = resp.Response
				}

				log.Printf("Last message: %s", lastMessage)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, messageReply)
				msg.ReplyToMessageID = lastMessageId

				_, err = bot.Send(msg)
				if nil != err {
					log.Fatalf("Error send msg: %s", err.Error())
				}
			}

			lastMessage = update.Message.Text
			lastMessageId = update.Message.MessageID
		}
	}
}
