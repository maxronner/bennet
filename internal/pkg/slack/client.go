package slack

import (
	"github.com/slack-go/slack"
)

func NewClient(token string, channelID string) *Client {
	return &Client{
		slack: slack.New(token),
		channelID: channelID,
	}
}

type Client struct {
	slack *slack.Client
	channelID string
}

// SendMessage sends a message via Slack to general channel
func (c *Client) SendMessage(message string) error {
	attachment := slack.Attachment {
		Pretext: message,
	}
	_, _, err := c.slack.PostMessage(
		c.channelID,
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true),
	)
	if err != nil {
		return err
	}
	return nil
}