package service_test

import (
	"github.com/nahuelPascual/tweeter/src/domain"
	"github.com/nahuelPascual/tweeter/src/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPublishedTweetsIsSaved(t *testing.T) {

	var tweet *domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	service.PublishTweet(tweet)

	assert.Equal(t, tweet, service.GetTweet(), "Expected tweet is")

}

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	var tweet *domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	// Operation
	service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweet()
	assert.Equal(t, text, publishedTweet.Text, "Expected tweetMsg is \"%s\" \nbut is \"%s\"", text, publishedTweet.Text)
	assert.Equal(t, user, publishedTweet.User, "Expected user is \"%s\" \nbut is \"%s\"", user, publishedTweet.User)
	assert.NotNil(t, publishedTweet.Date, "Expected date can't be nil")
}
