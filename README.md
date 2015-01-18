twi-list-slack
====

## Description
twitter list send to slack for golang

## Demo
![demo](https://qiita-image-store.s3.amazonaws.com/0/40887/2c919683-8cba-116d-4563-bfb57573a80a.png "スクリーンショット_2015-01-16_22_11_08.png")

## Usage

### config

Multiple slack channel support.

```bash
$ twi-list-slack config --filePath config.json
```

`config.json` is [sample](https://github.com/kyokomi/twi-list-slack/blob/master/config/config.json)

```
{
  "twitter": {
    "consumerKey": "xxxxxxxxxxxxxxxxxxxxx",
    "consumerSecret": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    "accessToken": "xxxxxxxxxxxxxx-xxxxxxxxxxxxxxxxxxxx",
    "accessTokenSecret": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  },
  "filters": [
    {
      "channelName": "#twitter",
      "incomingURL": "https://hooks.slack.com/services/xxxxxxxxxx/xxxxxxxxxxxxxxxxx/xxxxxxxxxxxxx"
    },
    {
      "channelName": "#random",
      "incomingURL": "https://hooks.slack.com/services/xxxxxx/xxxx/xxxxxx",
      "list_id": "111111111"
    }
  ]
}
```

### 'env' or 'flag'

```bash
$ twi-list-slack
```

## Install

To install, use `go get`:

```bash
$ go get github.com/kyokomi/twi-list-slack
```

## Setup 
`env` or `flag`

### Env

```bash
# Twitter OAuth
export TWLS_CONSUMER_KEY=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
export TWLS_CONSUMER_SECRET=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
export TWLS_ACCESS_TOKEN=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
export TWLS_ACCESS_TOKEN_SECRET=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
# Twitter Filter ListID
export TWLS_LIST_ID=123456789
# Slack incoming URL
export TWLS_INCOMING_URL=https://hooks.slack.com/services/xxxxxxxxxxxxxxxxxxxxxxxxx
```

- [TWLS_LIST_ID](https://dev.twitter.com/rest/reference/get/lists/members): The numerical id of the list.
- [TWLS_INCOMING_URL](https://my.slack.com/services/new/incoming-webhook): Send data into Slack in real-time.

imcoming URL:

![imcoming1](https://qiita-image-store.s3.amazonaws.com/0/40887/f8338a8f-1189-8889-7e1f-7feb7c416bd6.png "スクリーンショット_2015-01-16_20_58_27.png")

![imcoming2](https://qiita-image-store.s3.amazonaws.com/0/40887/6013f80f-c44b-3ac7-100d-acc03e9df447.png "スクリーンショット_2015-01-16_20_59_01.png")

### Flags

```bash
$ twi-list-slack  --help
NAME:
   twi-list-slack -

USAGE:
   twi-list-slack [global options] command [command options] [arguments...]

VERSION:
   0.1.0

AUTHOR:
  kyokomi - <kyoko1220adword@gmail.com>

COMMANDS:
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --ck 		twitter Consumer key [$TWLS_CONSUMER_KEY]
   --cs 		twitter Consumer secret [$TWLS_CONSUMER_SECRET]
   --at 		twitter Access token [$TWLS_ACCESS_TOKEN]
   --ats 		twitter Access token secret [$TWLS_ACCESS_TOKEN_SECRET]
   --list-id 		the numerical id of the list [$TWLS_LIST_ID]
   --cn '#twitter'	slack channel name [$TWLS_SLACK_CHANNEL_NAME]
   --incomingURL 	slack incomingURL [$TWLS_INCOMING_URL]
   --help, -h		show help
   --version, -v	print the version
```

## Contribution

1. Fork ([https://github.com/kyokomi/twi-list-slack/fork](https://github.com/kyokomi/twi-list-slack/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create new Pull Request

## Author

[kyokomi](https://github.com/kyokomi)
