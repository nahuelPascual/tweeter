package service

import (
	"fmt"
	"github.com/nahuelPascual/tweeter/src/domain"
	"strings"
)

type TweetManager struct {
	allTweets    []*domain.Tweet
	tweetsByUser map[string][]*domain.Tweet
}

func NewTweetManager() (manager *TweetManager) {
	manager = new(TweetManager)
	manager.tweetsByUser = make(map[string][]*domain.Tweet)
	return
}

func (manager *TweetManager) PublishTweet(tweety *domain.Tweet) (id int64, err error) {

	length := len(strings.TrimSpace(tweety.Text))

	if tweety.User == "" {
		err = fmt.Errorf("username is required")
	} else if length == 0 {
		err = fmt.Errorf("text is required")
	} else if length > 140 {
		err = fmt.Errorf("text limit is 140 characters")
	} else {
		manager.addNewTweet(tweety)
	}

	return tweety.Id, err
}

func (manager *TweetManager) GetTweet() *domain.Tweet {
	return manager.allTweets[len(manager.allTweets)-1]
}

func (manager *TweetManager) GetTweets() []*domain.Tweet {
	return manager.allTweets
}

func (manager *TweetManager) GetTweetById(id int64) (aTweet *domain.Tweet, err error) {
	for i := 0; i < len(manager.allTweets); i++ {
		if manager.allTweets[i].Id == id {
			aTweet = manager.allTweets[i]
			break
		}
	}
	if aTweet == nil {
		err = fmt.Errorf("el Tweet con id #%d no existe", id)
	}
	return aTweet, err
}

func (manager *TweetManager) CountTweetsByUser(username string) int {
	return len(manager.tweetsByUser[username])
}

func (manager *TweetManager) GetTweetsByUser(username string) (tweets []*domain.Tweet) {
	tweets = manager.tweetsByUser[username]
	return
}

func (manager *TweetManager) addNewTweet(tweet *domain.Tweet) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Print("Error!", r)
		}
	}()

	/* Adding tweet to allTweets slice */
	manager.allTweets = append(manager.allTweets, tweet)

	/* Adding tweet to map tweetsByUser */
	manager.tweetsByUser[tweet.User] = append(manager.tweetsByUser[tweet.User], tweet)
}
