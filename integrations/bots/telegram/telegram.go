package telegrambot

import (
	"fmt"
)

/*
the package defines a struct `TelegramClient` that holds the Telegram bot API and it also defines methods `SendMessage`, `EditMessage`, `DeleteMessage`, `AddReaction` and `GetUpdates` that can be used to interact with Telegram.

`NewTelegramClient` function creates and returns a new TelegramClient, it takes the token of the bot as a parameter and creates a new bot using the `tgbotapi.NewBotAPI(token)` function.

`SendMessage` function sends a message to a chat using the chatID

*/
// TelegramClient is a struct that holds the Telegram bot API
type TelegramClient struct {
	bot *tgbotapi.BotAPI
}

// NewTelegramClient creates and returns a new TelegramClient
func NewTelegramClient(token string) (*TelegramClient, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("Error creating Telegram bot: %s", err)

	}
	return &TelegramClient{bot: bot}, nil

}

// SendMessage sends a message to a chat
func (t *TelegramClient) SendMessage(chatID int64, message string) (int, error) {
	msg := tgbotapi.NewMessage(chatID, message)
	message, err := t.bot.Send(msg)
	if err != nil {
		return 0, fmt.Errorf("Error sending message: %s", err)

	}
	return message.MessageID, nil

}

// EditMessage edit a message in a chat
func (t *TelegramClient) EditMessage(chatID int64, messageID int, newMessage string) error {
	msg := tgbotapi.NewEditMessageText(chatID, messageID, newMessage)
	_, err := t.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("Error editing message: %s", err)

	}
	return nil

}

// DeleteMessage delete a message in a chat
func (t *TelegramClient) DeleteMessage(chatID int64, messageID int) error {
	msg := tgbotapi.NewDeleteMessage(chatID, messageID)
	_, err := t.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("Error deleting message: %s", err)

	}
	return nil
}

// AddReaction adds a reaction to a message
func (t *TelegramClient) AddReaction(chatID int64, messageID int, emoji string) error {
	msg := tgbotapi.NewMessage(chatID, emoji)
	msg.ReplyToMessageID = messageID
	_, err := t.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("Error adding reaction: %s", err)

	}
	return nil

}

// GetUpdates gets updates from Telegram
func (t *TelegramClient) GetUpdates() ([]tgbotapi.Update, error) {
	updates, err := t.bot.GetUpdates()
	if err != nil {
		return nil, fmt.Errorf("Error getting updates: %s", err)

	}
	return updates, nil

}
