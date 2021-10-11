package rest

import "encoding/json"

// SendMessage sends a message via Slack to general channel
func (s *Server) SendMessage(body []byte) error {
	var params struct {
		Message string `json:"message,omitempty"`
	}
	if err := json.Unmarshal(body, &params); err != nil {
		return err
	}
	if err := s.service.SendMessage(params.Message); err != nil {
		return err
	}
	return nil
}