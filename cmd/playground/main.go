package main

import (
	"os"

	"github.com/maxronner/bennet/internal/pkg/slack"
)

func main() {
	client := slack.NewClient(os.Getenv("TOKEN"), os.Getenv("CHANNEL_ID"))
	client.SendMessage(os.Args[1])
}