package service

var tweet string

func PublishTweet(tweetMsg string) {

	tweet = tweetMsg

}

func GetTweet() string {
	return tweet
}
