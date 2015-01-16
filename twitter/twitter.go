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

// Client is Twitter Client
type Client struct {
	consumer *oauth.Consumer
	token    *oauth.AccessToken

	Lists *ListsService
	User  *UserService
}

// NewClient create Twitter Client
func NewClient(consumerKey, consumerSecret, accessToken, accessTokenSecret string) *Client {
	c := Client{}

	sp := oauth.ServiceProvider{
		RequestTokenUrl:   requestTokenURL,
		AuthorizeTokenUrl: authorizeTokenURL,
		AccessTokenUrl:    accessTokenURL,
	}
	c.consumer = oauth.NewConsumer(consumerKey, consumerSecret, sp)

	c.token = &oauth.AccessToken{
		Token:  accessToken,
		Secret: accessTokenSecret,
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
