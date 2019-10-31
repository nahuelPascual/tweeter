package domain

import "time"

type Tweet struct {
	User string
	Text string
	Date *time.Time
}

func NewTweet(user, msg string) *Tweet {
	now := time.Now()
	return &Tweet{user, msg, &now}
}
