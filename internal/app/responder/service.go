package responder

import "github.com/maxronner/bennet/internal/pkg/slack"

func NewService(slack *slack.Client) *Service {
	return &Service{
		slack,
	}
}

type Service struct {
	slack *slack.Client
}

// SendMessage sends a message via Slack to general channel
func (s *Service) SendMessage(message string) error {
	if err:= s.slack.SendMessage(message); err != nil {
		return err
	}
	return nil
}