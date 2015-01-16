package main

import (
	"os"

	"github.com/codegangsta/cli"

	"fmt"
	"log"

	"encoding/json"

	"github.com/kyokomi/twi-list-slack/slack"
	"github.com/kyokomi/twi-list-slack/twitter"
	"github.com/ttacon/chalk"
)

func main() {
	app := cli.NewApp()
	app.Name = "twi-list-slack"
	app.Version = Version
	app.Usage = ""
	app.Author = "kyokomi"
	app.Email = "kyoko1220adword@gmail.com"
	app.Action = doMain
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "ck",
			Value:  "",
			Usage:  "twitter Consumer key",
			EnvVar: "TWLS_CONSUMER_KEY",
		},
		cli.StringFlag{
			Name:   "cs",
			Value:  "",
			Usage:  "twitter Consumer secret",
			EnvVar: "TWLS_CONSUMER_SECRET",
		},
		cli.StringFlag{
			Name:   "at",
			Value:  "",
			Usage:  "twitter Access token",
			EnvVar: "TWLS_ACCESS_TOKEN",
		},
		cli.StringFlag{
			Name:   "ats",
			Value:  "",
			Usage:  "twitter Access token secret",
			EnvVar: "TWLS_ACCESS_TOKEN_SECRET",
		},
		cli.StringFlag{
			Name:   "list-id",
			Value:  "",
			Usage:  "the numerical id of the list",
			EnvVar: "TWLS_LIST_ID",
		},
		cli.StringFlag{
			Name:   "cn",
			Value:  "#twitter",
			Usage:  "slack channel name",
			EnvVar: "TWLS_SLACK_CHANNEL_NAME",
		},
		cli.StringFlag{
			Name:   "incomingURL",
			Value:  "",
			Usage:  "slack incomingURL",
			EnvVar: "TWLS_INCOMING_URL",
		},
	}
	app.Run(os.Args)
}

func doMain(c *cli.Context) {

	// TwitterClient
	ck := c.GlobalString("ck")
	cs := c.String("cs")
	at := c.String("at")
	ats := c.String("ats")
	channelName := c.String("cn")

	var tc *twitter.Client
	tc = twitter.NewClient(ck, cs, at, ats)

	// SlackClient
	incomingURL := c.String("incomingURL")

	var sc *slack.Client
	sc = slack.NewClient(incomingURL)

	// Filter生成
	// TODO: とりあえず一旦はListID固定 設定ファイルとかから読み込むようにしたい
	listID := c.String("list-id")
	filter, err := NewListIDFilter(tc, listID)
	if err != nil {
		log.Println(err)
	}

	// Streamingする
	params := make(map[string]string)
	params["replies"] = "all"
	tc.User.Stream(params, func(data []byte) bool {
		text := string(data)
		if len(text) == 0 {
			return false
		}

		var s twitter.Streaming
		if err := json.Unmarshal(data, &s); err != nil {
			fmt.Println("error json parse => ", text)
			// TODO: favったらeventとか別のjsonがくるから一旦logだけ出してスルー
			//			return true
			return false
		}

		if !filter.filter(s) {
			return false
		}

		// TODO: debug log
		h := chalk.Yellow.Color(fmt.Sprintf("%s : %s@%s \n", s.CreatedAt, s.User.Name, s.User.ScreenName))
		b := fmt.Sprintf("> %s\n", s.Text)
		fmt.Print(h, b)

		// Slack Send
		go func() {
			message := slack.OutgoingMessage{}
			message.Channel = channelName
			message.Username = fmt.Sprintf("%s@%s", s.User.Name, s.User.ScreenName)
			message.Text = s.Text
			message.IconURL = s.User.ProfileImageURL

			if err := sc.SendMessage(message); err != nil {
				fmt.Println("error send message => ", err)
			}
		}()

		return false
	})
}
