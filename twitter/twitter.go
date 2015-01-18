package twitter

import (
	"net/http"

	"github.com/mrjones/oauth"
)

const (
	requestTokenURL   = "http://api.twitter.com/oauth/request_token"
	authorizeTokenURL = "https://api.twitter.com/oauth/authorize"
	accessTokenURL    = "https://api.twitter.com/oauth/access_token"
)

type Config struct {
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
}

// Client is Twitter Client
type Client struct {
	consumer *oauth.Consumer
	token    *oauth.AccessToken

	Lists *ListsService
	User  *UserService
}

// NewClient create Twitter Client
func NewClient(consumerKey, consumerSecret, accessToken, accessTokenSecret string) *Client {
	c := Config{}
	c.ConsumerKey = consumerKey
	c.ConsumerSecret = consumerSecret
	c.AccessToken = accessToken
	c.AccessTokenSecret = accessTokenSecret
	return NewClientConfig(c)
}

// NewClient create Twitter Client
func NewClientConfig(config Config) *Client {
	c := Client{}

	sp := oauth.ServiceProvider{
		RequestTokenUrl:   requestTokenURL,
		AuthorizeTokenUrl: authorizeTokenURL,
		AccessTokenUrl:    accessTokenURL,
	}
	c.consumer = oauth.NewConsumer(config.ConsumerKey, config.ConsumerSecret, sp)

	c.token = &oauth.AccessToken{
		Token:  config.AccessToken,
		Secret: config.AccessTokenSecret,
	}

	c.Lists = &ListsService{client: &c}
	c.User = &UserService{client: &c}

	return &c
}

func (c *Client) Get(url string, userParams map[string]string) (resp *http.Response, err error) {
	return c.consumer.Get(url, userParams, c.token)
}

func (c *Client) Post(url string, userParams map[string]string) (resp *http.Response, err error) {
	return c.consumer.Post(url, userParams, c.token)
}

func (c *Client) Delete(url string, userParams map[string]string) (resp *http.Response, err error) {
	return c.consumer.Delete(url, userParams, c.token)
}

func (c *Client) Put(url string, body string, userParams map[string]string) (resp *http.Response, err error) {
	return c.consumer.Put(url, body, userParams, c.token)
}
