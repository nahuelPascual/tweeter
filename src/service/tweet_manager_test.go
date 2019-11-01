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

	_, _ = service.PublishTweet(tweet)

	assert.Equal(t, tweet, service.GetTweet(), "Expected tweet is")

}

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	var tweet *domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	// Operation
	_, _ = service.PublishTweet(tweet)

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
	_, err = service.PublishTweet(tweet)

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
	_, err = service.PublishTweet(tweet)

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
	_, err = service.PublishTweet(tweet)

	// Validation
	assert.Error(t, err)
	assert.Equal(t, "text limit is 140 characters", err.Error(), "Expected error is 'text limit is 140 characters'")
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {

	// Initialization
	service.InitializeService()
	var tweet, secondTweet *domain.Tweet
	tweet = domain.NewTweet("npascual", "First tweet")
	secondTweet = domain.NewTweet("npascual", "Second tweet")

	// Operation
	_, _ = service.PublishTweet(tweet)
	_, _ = service.PublishTweet(secondTweet)

	// Validation
	publishedTweets := service.GetTweets()
	assert.Equal(t, 2, len(publishedTweets), "Expected size is 2 but was %d", len(publishedTweets))
}

func TestCanRetrieveTweetById(t *testing.T) {

	// Initialization
	service.InitializeService()

	var tweet *domain.Tweet
	var id int64

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	id, _ = service.PublishTweet(tweet)

	// Validation
	publishedTweet, _ := service.GetTweetById(id)

	assert.Equal(t, id, publishedTweet.Id, "Expected id #%d", id)
}

func TestCannotRetrieveTweetByInexistentId(t *testing.T) {

	// Initialization
	service.InitializeService()

	var id int64 = 9777797979

	// Validation
	_, err := service.GetTweetById(id)

	assert.Error(t, err, "Expected id #%d", id)
}
