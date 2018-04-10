package notify

import "testing"

func TestBasic(t *testing.T) {
	test := func(n Notifier) {}
	test(NewTwitter(TwitterConfig{}))
	test(NewSlack("https://slack.com"))
}
