package service

import (
	"fmt"
	"github.com/nahuelPascual/tweeter/src/domain"
	"strings"
)

var allTweets []*domain.Tweet
var tweetsByUser = map[string][]*domain.Tweet{}

func PublishTweet(tweety *domain.Tweet) (id int64, err error) {

	length := len(strings.TrimSpace(tweety.Text))

	if tweety.User == "" {
		err = fmt.Errorf("username is required")
	} else if length == 0 {
		err = fmt.Errorf("text is required")
	} else if length > 140 {
		err = fmt.Errorf("text limit is 140 characters")
	} else {
		addNewTweet(tweety)
	}

	return tweety.Id, err
}

func GetTweet() *domain.Tweet {
	return allTweets[len(allTweets)-1]
}

func GetTweets() []*domain.Tweet {
	return allTweets
}

func InitializeService() {
	allTweets = make([]*domain.Tweet, 0)
	tweetsByUser = make(map[string][]*domain.Tweet)
}

func GetTweetById(id int64) (aTweet *domain.Tweet, err error) {
	for i := 0; i < len(allTweets); i++ {
		if allTweets[i].Id == id {
			aTweet = allTweets[i]
			break
		}
	}
	if aTweet == nil {
		err = fmt.Errorf("el Tweet con id #%d no existe", id)
	}
	return aTweet, err
}

func CountTweetsByUser(username string) int {
	return len(tweetsByUser[username])
}

func addNewTweet(tweet *domain.Tweet) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Print("Error!", r)
		}
	}()

	/* Adding tweet to allTweets slice */
	allTweets = append(allTweets, tweet)

	/* Adding tweet to map tweetsByUser */
	tweetsByUser[tweet.User] = append(tweetsByUser[tweet.User], tweet)
}
