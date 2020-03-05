package main

import (
	"log"
	"os"

	"github.com/alienvspredator/telegram-bot/lib/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func anyCondition(message telegram.Message) bool {
	return true
}

func logFunc(message telegram.Message, botAPI *telegram.BotAPI, bot *telegram.Bot, next func()) {
	log.Printf("[id%d] %s %s: %s\n", message.Chat.ID, message.Chat.FirstName, message.Chat.LastName, message.Text)
	next()
}

func stopMiddleware(message telegram.Message, botAPI *telegram.BotAPI, bot *telegram.Bot, next func()) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Прощай, жестокий мир!")
	botAPI.Send(msg)
	bot.Stop()
}

func bagBonesMiddleware(message telegram.Message, botAPI *telegram.BotAPI, bot *telegram.Bot, next func()) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Чего надобно, мешок с костями?")
	botAPI.Send(msg)
}

func unhandledMiddleware(message telegram.Message, botAPI *telegram.BotAPI, bot *telegram.Bot, next func()) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Не знаю что ты мне пишешь, но зовут тебя "+message.Chat.FirstName)
	botAPI.Send(msg)
}

func main() {
	token := os.Getenv("TOKEN")
	botParams := &telegram.BotParams{Token: token}
	bot, err := telegram.NewBot(botParams)
	if err != nil {
		panic(err)
	}

	bot.OnCondition(anyCondition, logFunc)
	bot.OnMessage("/kill", stopMiddleware)
	bot.OnMessage("/start", bagBonesMiddleware)
	bot.OnCondition(anyCondition, unhandledMiddleware)
	bot.Start()
}
