package main

import (
	"log"
	"os"

	"github.com/alienvspredator/telegram-bot/lib/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func onHello(message telegram.Message, botAPI *telegram.BotAPI, bot *telegram.Bot) {
	answer := tgbotapi.NewMessage(message.Chat.ID, "Hi!")
	botAPI.Send(answer)
}

func anyCondition(message telegram.Message) bool {
	return true
}

func logFunc(message telegram.Message, botAPI *telegram.BotAPI, bot *telegram.Bot) {
	log.Printf("[id%d] %s %s: %s\n", message.Chat.ID, message.Chat.FirstName, message.Chat.LastName, message.Text)
}

func stopCondition(message telegram.Message) bool {
	return message.Text == "/stop"
}

func stop(message telegram.Message, botAPI *telegram.BotAPI, bot *telegram.Bot) {
	bot.Stop()
}

func main() {
	token := os.Getenv("TOKEN")
	botParams := &telegram.BotParams{Token: token}
	bot, err := telegram.NewBot(botParams)
	if err != nil {
		panic(err)
	}

	bot.OnCondition(anyCondition, logFunc)
	bot.OnCondition(stopCondition, stop)
	bot.OnMessage("Hello", onHello)
	bot.Start()
}
