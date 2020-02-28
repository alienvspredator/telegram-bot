package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func mkBot(botAPI *tgbotapi.BotAPI) *Bot {
	return &Bot{botAPI: botAPI, conditions: make(map[*func(message tgbotapi.Message) bool][]func(message tgbotapi.Message, botAPI *tgbotapi.BotAPI, bot *Bot, next func())), sigTerm: -1}
}

// NewBot creates new bot
func NewBot(params *BotParams) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(params.Token)
	if err != nil {
		return nil, err
	}

	return mkBot(bot), nil
}
