package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// OnMessageCallback is callback of message event
type OnMessageCallback = func(message tgbotapi.Message, botAPI *tgbotapi.BotAPI, bot *Bot, next func())

// Condition function checks is the message satisfies to be called
type Condition = func(message tgbotapi.Message) bool

// Message contains data about almost anything.
type Message = tgbotapi.Message

// BotAPI allows you to interact with the Telegram Bot API
type BotAPI = tgbotapi.BotAPI

// Bot is struct of bot
type Bot struct {
	conditions map[*Condition][]OnMessageCallback
	botAPI     *tgbotapi.BotAPI
	sigTerm    int
}

// BotParams is structure with params of the bot
type BotParams struct {
	Token string `json:"token"`
}
