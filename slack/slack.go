package slack

import (
	"net/http"
	"net/url"

	"encoding/json"
)

type Client struct {
	incomingURL string
}

func NewClient(incomingURL string) *Client {
	c := Client{}

	c.incomingURL = incomingURL

	return &c
}

// OutgoingMessage is Slack PostRequest Message.
type OutgoingMessage struct {
	IconURL  string `json:"icon_url"`
	Channel  string `json:"channel"`
	Username string `json:"username"`
	Text     string `json:"text"`
}

func (c Client) SendMessage(message OutgoingMessage) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	if _, err := http.PostForm(c.incomingURL, url.Values{"payload": {string(body)}}); err != nil {
		return err
	}

	return nil
}
