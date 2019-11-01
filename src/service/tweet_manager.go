package service

import (
	"fmt"
	"github.com/nahuelPascual/tweeter/src/domain"
	"strings"
)

var tweets []*domain.Tweet

func PublishTweet(tweety *domain.Tweet) (id int64, err error) {

	length := len(strings.TrimSpace(tweety.Text))

	if tweety.User == "" {
		err = fmt.Errorf("username is required")
	} else if length == 0 {
		err = fmt.Errorf("text is required")
	} else if length > 140 {
		err = fmt.Errorf("text limit is 140 characters")
	} else {
		tweets = append(tweets, tweety)
	}

	return tweety.Id, err
}

func GetTweet() *domain.Tweet {
	return tweets[len(tweets)-1]
}

func GetTweets() []*domain.Tweet {
	return tweets
}

func InitializeService() {
	tweets = make([]*domain.Tweet, 0)
}

func GetTweetById(id int64) (aTweet *domain.Tweet, err error) {
	for i := 0; i < len(tweets); i++ {
		if tweets[i].Id == id {
			aTweet = tweets[i]
			break
		}
	}
	if aTweet == nil {
		err = fmt.Errorf("El Tweet con id #%d no existe", id)
	}
	return aTweet, err
}
