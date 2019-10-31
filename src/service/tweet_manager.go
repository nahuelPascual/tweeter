package service

import "github.com/nahuelPascual/tweeter/src/domain"

var tweet *domain.Tweet

func PublishTweet(tweetMsg *domain.Tweet) {

	tweet = tweetMsg

}

func GetTweet() *domain.Tweet {
	return tweet
}
