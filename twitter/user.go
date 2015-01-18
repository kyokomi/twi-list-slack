package twitter

import "bufio"

const (
	userStreamURL = "https://userstream.twitter.com/1.1/user.json"
)

type UserService struct {
	client *Client
}

type Handle func(data []byte) bool

func (u *UserService) Stream(params map[string]string, streamHandler Handle) error {

	res, err := u.client.Get(userStreamURL, params)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	scan := bufio.NewScanner(res.Body)

	if scan.Scan() {
		// TODO: 1回目はフレンド一覧
	}

	for scan.Scan() {
		if streamHandler(scan.Bytes()) {
			break
		}
	}

	return nil
}
