package domain_test

import (
	"github.com/nahuelPascual/tweeter/src/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanGetAPrintableTweet(t *testing.T) {

	// Initialization
	tweet := domain.NewTweet("grupoesfera", "This is my tweet")

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	assert.Equal(t, text, expectedText, "The expected text is %s but was %s", expectedText, text)

}
