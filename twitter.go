package notify

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Twitter struct {
	*twitter.Client
}

type TwitterConfig struct {
	ConsumerKey    string `json:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret"`
	AccessToken    string `json:"access_token"`
	AccessSecret   string `json:"access_secret"`
}

func NewTwitter(c TwitterConfig) *Twitter {
	config := oauth1.NewConfig(c.ConsumerKey, c.ConsumerSecret)
	token := oauth1.NewToken(c.AccessToken, c.AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	return &Twitter{twitter.NewClient(httpClient)}
}

func (t *Twitter) Notify(text string) error {
	_, _, err := t.Statuses.Update(text, nil)
	return err
}
