package telegram

import (
	"reflect"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func messageCondition(message tgbotapi.Message, messageText string) bool {
	return reflect.TypeOf(message.Text).Kind() == reflect.String && message.Text == messageText
}

// OnMessage registers callback on message
func (bot *Bot) OnMessage(messageText string, callback OnMessageCallback) {
	bot.OnCondition(func(message tgbotapi.Message) bool { return messageCondition(message, messageText) }, callback)
}

// OnCondition registers callback on any message that satisfies condition
func (bot *Bot) OnCondition(condition Condition, callback OnMessageCallback) {
	if bot.conditions[&condition] == nil {
		bot.conditions[&condition] = make([]func(message tgbotapi.Message, botAPI *tgbotapi.BotAPI, bot *Bot, next func()), 0)
	}

	bot.conditions[&condition] = append(bot.conditions[&condition], callback)
}
