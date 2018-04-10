package notify

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Slack struct {
	URL       string
	Channel   string
	IconEmoji string
	Username  string
}

func NewSlack(URL string) *Slack {
	return &Slack{
		URL:       URL,
		Channel:   "#general",
		IconEmoji: ":penguin:",
		Username:  "golang-app",
	}
}

type payload struct {
	Channel   string `json:"channel"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	Text      string `json:"text"`
}

func (s *Slack) Notify(text string) error {
	p := payload{
		Channel:   s.Channel,
		Username:  s.Username,
		IconEmoji: s.IconEmoji,
		Text:      text,
	}
	b, err := json.Marshal(p)
	if err != nil {
		return err
	}

	resp, err := http.Post(s.URL, "application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode/100 == 2 {
		return nil
	} else {
		return errors.New(resp.Status)
	}
}
