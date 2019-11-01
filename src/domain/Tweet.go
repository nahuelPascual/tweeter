package domain

import (
	"time"
)

var seq_id int64 = 1

type Tweet struct {
	Id   int64
	User string
	Text string
	Date *time.Time
}

func NewTweet(user, msg string) *Tweet {
	now := time.Now()
	id := seq_id
	seq_id += 1
	return &Tweet{id, user, msg, &now}
}
