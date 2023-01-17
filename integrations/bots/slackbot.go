package slackbot

import (
	"fmt"

	"github.com/nlopes/slack"
)

/*
the package defines a struct SlackClient that holds the slack API client, and it also defines three methods PostMessage, GetConversationHistory, and AddReaction that can be used to interact with Slack.
*/
// SlackClient is a struct that holds the slack API client
type SlackClient struct {
	api *slack.Client
}

// NewSlackClient creates and returns a new SlackClient
func NewSlackClient(token string) *SlackClient {
	return &SlackClient{api: slack.New(token)}

}

// PostMessage sends a message to a channel
func (s *SlackClient) PostMessage(channel, message string) (string, string, error) {
	channelID, timestamp, err := s.api.PostMessage(channel, slack.MsgOptionText(message, false))
	if err != nil {
		return "", "", fmt.Errorf("Error sending message: %s", err)

	}
	return channelID, timestamp, nil

}

// EditMessage edit a message in a channel
func (s *SlackClient) EditMessage(channel, timestamp, message string) error {
	_, _, err := s.api.UpdateMessage(channel, timestamp, slack.MsgOptionText(message, false))
	if err != nil {
		return fmt.Errorf("Error editing message: %s", err)

	}
	return nil

}

// DeleteMessage delete a message in a channel
func (s *SlackClient) DeleteMessage(channel, timestamp string) error {
	err := s.api.DeleteMessage(channel, timestamp)
	if err != nil {
		return fmt.Errorf("Error deleting message: %s", err)

	}
	return nil

}

// GetConversationHistory returns the conversation history of a channel
func (s *SlackClient) GetConversationHistory(channel string) ([]slack.Message, error) {
	params := slack.GetConversationHistoryParameters{}
	history, err := s.api.GetConversationHistory(channel, params)
	if err != nil {
		return nil, fmt.Errorf("Error getting conversation history: %s", err)

	}
	return history.Messages, nil

}

// AddReaction adds a reaction to a message
func (s *SlackClient) AddReaction(emoji, messageRef string) error {
	err := s.api.AddReaction(emoji, slack.NewRefToMessage(messageRef))
	if err != nil {
		return fmt.Errorf("Error adding reaction: %s", err)

	}
	return nil

}
