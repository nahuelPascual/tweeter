package service

import (
	"fmt"
	"github.com/nahuelPascual/tweeter/src/domain"
	"strings"
)

var tweet *domain.Tweet

func PublishTweet(tweety *domain.Tweet) (err error) {

	length := len(strings.TrimSpace(tweety.Text))

	if tweety.User == "" {
		err = fmt.Errorf("username is required")
	} else if length == 0 {
		err = fmt.Errorf("text is required")
	} else if length > 140 {
		err = fmt.Errorf("text limit is 140 characters")
	} else {
		tweet = tweety
	}

	return err
}

func GetTweet() *domain.Tweet {
	return tweet
}
