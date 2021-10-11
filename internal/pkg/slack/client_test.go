//go:build integration

package slack_test

import (
	"testing"
	"time"

	"github.com/maxronner/bennet/internal/pkg/slack"
)

func TestClient_SendMessage(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name    string
		c       *slack.Client
		args    args
		wantErr bool
	}{
		{ 
			name: "bad token",
			c: slack.NewClient("TEST_CHANNEL_ID","blaj"),
			args: args{
				message: "Test",
			},
			wantErr: true,
		},
		{ 
			name: "send on general",
			c: slack.NewTestClient(),
			args: args{
				message: time.Now().Format("2006-01-02 15:04:05"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.SendMessage(tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("Client.SendMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}