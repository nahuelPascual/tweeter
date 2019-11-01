package main

import (
	"github.com/abiosoft/ishell"
	"github.com/nahuelPascual/tweeter/src/domain"
	"github.com/nahuelPascual/tweeter/src/service"
	"strconv"
)

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	tweetManager := service.NewTweetManager()

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your username: ")
			user := c.ReadLine()

			c.Print("Write your tweet: ")
			tweetMsg := c.ReadLine()

			tweet := domain.NewTweet(user, tweetMsg)

			_, err := tweetManager.PublishTweet(tweet)

			if err == nil {
				c.Print("Tweet sent\n")
			} else {
				c.Print("Error!\n" + err.Error() + "\n")
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showLastTweet",
		Help: "Shows the last tweet published",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := tweetManager.GetTweet()

			c.Println(*tweet)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweetById",
		Help: "shows the tweet that matches with id provided",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			id, _ := strconv.ParseInt(c.Args[0], 10, 64)

			tweet, er := tweetManager.GetTweetById(id)

			if er != nil {
				c.Println(er)
			} else {
				c.Print(tweet)
			}
		},
	})

	shell.Run()

}
