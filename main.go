package main

import (
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"net/http"
	"os"
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

	//Health checker
	go func() {
		router := gin.Default()
		router.GET("/", func(context *gin.Context) {
			context.String(http.StatusOK, "Timeweb Cloud + Gin = ❤️")
		})
		err := router.Run()
		if err != nil {
			return
		}
	}()

	for update := range updates {
		if update.Message != nil {
			if update.Message.Text == triggerMessage {
				messageReply := "Брат, чет отвалилось, не смогу получить ответ"
				resp, err := chadGpt.GetAnswer(lastMessage)
				if nil == err {
					messageReply = resp.Response
				}

				if len(messageReply) < 1 {
					messageReply = "Вы меня задрочили и я не могу получить ответ от Chat Gpt"
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
