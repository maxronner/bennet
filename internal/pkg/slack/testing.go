package slack

import "os"

func NewTestClient() *Client {
	return NewClient(os.Getenv("TEST_TOKEN"), os.Getenv("TEST_CHANNEL_ID"))
}