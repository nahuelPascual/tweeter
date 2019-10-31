package service_test

import (
	"github.com/nahuelPascual/tweeter/src/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPublishedTweetsIsSaved(t *testing.T) {

	tweet := "This is my first tweet"

	service.PublishTweet(tweet)

	assert.Equal(t, tweet, service.GetTweet(), "Expected tweet is")

}
