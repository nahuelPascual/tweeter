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

	_ = service.PublishTweet(tweet)

	assert.Equal(t, tweet, service.GetTweet(), "Expected tweet is")

}

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	var tweet *domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	// Operation
	_ = service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweet()
	assert.Equal(t, text, publishedTweet.Text, "Expected tweetMsg is \"%s\" \nbut is \"%s\"", text, publishedTweet.Text)
	assert.Equal(t, user, publishedTweet.User, "Expected user is \"%s\" \nbut is \"%s\"", user, publishedTweet.User)
	assert.NotNil(t, publishedTweet.Date, "Expected date can't be nil")
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {

	// Initialization
	var tweet *domain.Tweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	err = service.PublishTweet(tweet)

	// Validation
	assert.Error(t, err)
	assert.Equal(t, "username is required", err.Error(), "Expected error is 'username is required'")
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

	// Initialization
	var tweet *domain.Tweet

	user := "npascual"
	var text string

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	err = service.PublishTweet(tweet)

	// Validation
	assert.Error(t, err)
	assert.Equal(t, "text is required", err.Error(), "Expected error is 'text is required'")
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {

	// Initialization
	var tweet *domain.Tweet

	user := "npascual"
	text := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	err = service.PublishTweet(tweet)

	// Validation
	assert.Error(t, err)
	assert.Equal(t, "text limit is 140 characters", err.Error(), "Expected error is 'text limit is 140 characters'")
}
