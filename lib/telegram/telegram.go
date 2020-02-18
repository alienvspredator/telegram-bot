package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Start creates and starts telegram bot
func (bot *Bot) Start() error {
	botAPI := bot.botAPI

	// Set update timeout
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// Get updates from bot
	updates, _ := botAPI.GetUpdatesChan(u)

	// Event loop
	for update := range updates {
		conditions := bot.conditions
		for condition, callbacks := range conditions {
			if (*condition)(*update.Message) {
				for _, callback := range callbacks {
					callback(*update.Message, bot.botAPI, bot)
				}
			}
		}

		if bot.sigTerm != -1 {
			botAPI.StopReceivingUpdates()
			return nil
		}
	}

	return nil
}

// Stop bot
func (bot *Bot) Stop() {
	bot.sigTerm = 0
}
