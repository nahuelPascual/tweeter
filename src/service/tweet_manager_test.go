package service_test

import (
	"github.com/nahuelPascual/tweeter/src/domain"
	"github.com/nahuelPascual/tweeter/src/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPublishedTweetsIsSaved(t *testing.T) {

	tweetManager := service.NewTweetManager()
	var tweet *domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	_, _ = tweetManager.PublishTweet(tweet)

	assert.Equal(t, tweet, tweetManager.GetTweet(), "Expected tweet is", tweet)

}

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()
	var tweet *domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	// Operation
	_, _ = tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := tweetManager.GetTweet()
	assert.Equal(t, text, publishedTweet.Text, "Expected tweetMsg is \"%s\" \nbut is \"%s\"", text, publishedTweet.Text)
	assert.Equal(t, user, publishedTweet.User, "Expected user is \"%s\" \nbut is \"%s\"", user, publishedTweet.User)
	assert.NotNil(t, publishedTweet.Date, "Expected date can't be nil")
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()
	var tweet *domain.Tweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	assert.Error(t, err)
	assert.Equal(t, "username is required", err.Error(), "Expected error is 'username is required'")
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()
	var tweet *domain.Tweet

	user := "npascual"
	var text string

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	assert.Error(t, err)
	assert.Equal(t, "text is required", err.Error(), "Expected error is 'text is required'")
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()
	var tweet *domain.Tweet

	user := "npascual"
	text := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	assert.Error(t, err)
	assert.Equal(t, "text limit is 140 characters", err.Error(), "Expected error is 'text limit is 140 characters'")
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()
	var tweet, secondTweet *domain.Tweet
	tweet = domain.NewTweet("npascual", "First tweet")
	secondTweet = domain.NewTweet("npascual", "Second tweet")

	// Operation
	_, _ = tweetManager.PublishTweet(tweet)
	_, _ = tweetManager.PublishTweet(secondTweet)

	// Validation
	publishedTweets := tweetManager.GetTweets()
	assert.Equal(t, 2, len(publishedTweets), "Expected size is 2 but was %d", len(publishedTweets))
}

func TestCanRetrieveTweetById(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet *domain.Tweet
	var id int64

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	id, _ = tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet, _ := tweetManager.GetTweetById(id)

	assert.Equal(t, id, publishedTweet.Id, "Expected id #%d", id)
}

func TestCannotRetrieveTweetByInexistentId(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var id int64 = 9777797979

	// Validation
	_, err := tweetManager.GetTweetById(id)

	assert.Error(t, err, "Expected id #%d", id)
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()
	var tweet, secondTweet, thirdTweet *domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"

	anotherUser := "nick"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)

	_, _ = tweetManager.PublishTweet(tweet)
	_, _ = tweetManager.PublishTweet(secondTweet)
	_, _ = tweetManager.PublishTweet(thirdTweet)

	// Operation
	count := tweetManager.CountTweetsByUser(user)

	// Validation
	assert.Equal(t, 2, count, "expected count is 2 but was %d", count)
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet := domain.NewTweet(user, text)
	secondTweet := domain.NewTweet(user, secondText)
	thirdTweet := domain.NewTweet(anotherUser, text)

	// Operation
	_, _ = tweetManager.PublishTweet(tweet)
	_, _ = tweetManager.PublishTweet(secondTweet)
	_, _ = tweetManager.PublishTweet(thirdTweet)
	tweets := tweetManager.GetTweetsByUser(user)

	// Validation
	assert.Equal(t, 2, len(tweets))
}
