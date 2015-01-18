package main

import (
	"os"

	"github.com/codegangsta/cli"

	"fmt"
	"log"

	"encoding/json"

	"github.com/kyokomi/twi-list-slack/config"
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
	app.Commands = []cli.Command{
		cli.Command{
			Name:      "config",
			ShortName: "c",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "filePath",
					Value: "./config/config.json",
					Usage: "config json file path",
				},
			},
			Action: doConfig,
		},
	}
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

type TwiListSlack struct {
	config       *config.TwiListSlackConfig
	twitter      *twitter.Client
	slackFilters []SlackFilter
}

type SlackFilter struct {
	StreamingFilter
	channelName string
	slack       *slack.Client
}

func doMain(c *cli.Context) {

	// TwitterClient
	ck := c.GlobalString("ck")
	cs := c.String("cs")
	at := c.String("at")
	ats := c.String("ats")
	channelName := c.String("cn")
	incomingURL := c.String("incomingURL")
	listID := c.String("list-id")

	conf := &config.TwiListSlackConfig{}
	conf.Twitter.ConsumerKey = ck
	conf.Twitter.ConsumerSecret = cs
	conf.Twitter.AccessToken = at
	conf.Twitter.AccessTokenSecret = ats
	conf.Filters = []config.Filter{
		config.Filter{
			IncomingURL: incomingURL,
			ChannelName: channelName,
			ListID:      listID,
		},
	}

	exec(conf)
}

func doConfig(c *cli.Context) {
	conf, err := config.NewConfig(c.String("filePath"))
	if err != nil {
		log.Fatalln(err)
	}

	exec(conf)
}

func exec(config *config.TwiListSlackConfig) {

	t := TwiListSlack{}

	t.config = config
	t.twitter = twitter.NewClientConfig(t.config.Twitter)

	t.slackFilters = make([]SlackFilter, 0)
	for _, filter := range t.config.Filters {

		var sc *slack.Client
		sc = slack.NewClient(filter.IncomingURL)

		listID := filter.ListID
		f, err := NewListIDFilter(t.twitter, listID)
		if err != nil {
			log.Println(err)
		}

		fmt.Printf("%s %+v\n", filter.ChannelName, f)

		sf := SlackFilter{
			StreamingFilter: f,
			channelName:     filter.ChannelName,
			slack:           sc,
		}
		t.slackFilters = append(t.slackFilters, sf)
	}

	t.stream()
}

func (t *TwiListSlack) stream() {

	// Streamingする
	params := make(map[string]string)
	params["replies"] = "all"
	t.twitter.User.Stream(params, func(data []byte) bool {
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

		for _, f := range t.slackFilters {
			if !f.StreamingFilter.filter(s) {
				return false
			}

			// TODO: debug log

			h := chalk.Yellow.Color(fmt.Sprintf("[%s] %s : %s@%s \n", f.channelName, s.CreatedAt, s.User.Name, s.User.ScreenName))
			b := fmt.Sprintf("> %s\n", s.Text)
			fmt.Print(h, b)

			// Slack Send
			message := slack.OutgoingMessage{}
			message.Channel = f.channelName
			message.Username = fmt.Sprintf("%s@%s", s.User.Name, s.User.ScreenName)
			message.Text = s.Text
			message.IconURL = s.User.ProfileImageURL

			if err := f.slack.SendMessage(message); err != nil {
				fmt.Println("error send message => ", err)
			}
		}

		return false
	})
}
