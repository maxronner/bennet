package main

import (
	"context"
	"os"

	"github.com/maxronner/bennet/internal/app/responder"
	responder_rest "github.com/maxronner/bennet/internal/app/responder/rest"
	"github.com/maxronner/bennet/internal/pkg/slack"
)

func main() {
	// Secondary adapter(s)
	slack := slack.NewClient(os.Getenv("TOKEN"), os.Getenv("CHANNEL_ID"))
	responderService := responder.NewService(slack)
	restResponderServer := responder_rest.NewServer(responderService)
	restResponderServer.Run(context.TODO())
}