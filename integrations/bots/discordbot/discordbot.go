package discordbot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

/*
the package defines a struct DiscordClient that holds the discord session and it also defines three methods SendMessage, EditMessage, DeleteMessage
*/
// DiscordClient is a struct that holds the discord session
type DiscordClient struct {
	session *discordgo.Session
}

// NewDiscordClient creates and returns a new DiscordClient
func NewDiscordClient(token string) (*DiscordClient, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("Error creating discord session: %s", err)

	}
	return &DiscordClient{session: session}, nil

}

// SendMessage sends a message to a channel
func (d *DiscordClient) SendMessage(channelID, message string) (string, error) {
	message, err := d.session.ChannelMessageSend(channelID, message)
	if err != nil {
		return "", fmt.Errorf("Error sending message: %s", err)

	}
	return message.ID, nil

}

// EditMessage edit a message in a channel
func (d *DiscordClient) EditMessage(channelID, messageID, newMessage string) error {
	_, err := d.session.ChannelMessageEdit(channelID, messageID, newMessage)
	if err != nil {
		return fmt.Errorf("Error editing message: %s", err)

	}
	return nil

}

// DeleteMessage delete a message in a channel
func (d *DiscordClient) DeleteMessage(channelID, messageID string) error {
	err := d.session.ChannelMessageDelete(channelID, messageID)
	if err != nil {
		return fmt.Errorf("Error deleting message: %s", err)

	}
	return nil

}

// AddReaction adds a reaction to a message
func (d *DiscordClient) AddReaction(channelID, messageID, emoji string) error {
	err := d.session.MessageReactionAdd(channelID, messageID, emoji)
	if err != nil {
		return fmt.Errorf("Error adding reaction: %s", err)

	}
	return nil

}
