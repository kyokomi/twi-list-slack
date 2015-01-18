package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/kyokomi/twi-list-slack/twitter"
)

type TwiListSlackConfig struct {
	Filters []Filter       `json:"filters"`
	Twitter twitter.Config `json:"twitter"`
}

type Filter struct {
	IncomingURL string `json:"incomingURL"`
	ChannelName string `json:"channelName"`
	ListID      string `json:"list_id"`
}

func NewConfig(fileName string) (*TwiListSlackConfig, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var c TwiListSlackConfig
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, err
	}

	return &c, nil
}
