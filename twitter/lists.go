package twitter

import (
	"encoding/json"
	"io/ioutil"
)

const (
	listsMembersURL = "https://api.twitter.com/1.1/lists/members.json"
)

type ListsService struct {
	client *Client
}

// ListsMembers list members
type ListsMembers struct {
	NextCursor        int    `json:"next_cursor"`
	NextCursorStr     string `json:"next_cursor_str"`
	PreviousCursor    int    `json:"previous_cursor"`
	PreviousCursorStr string `json:"previous_cursor_str"`
	Users             []User `json:"users"`
}

func (l *ListsService) GetMembers(listID string) (*ListsMembers, error) {

	params := make(map[string]string)
	params["list_id"] = listID

	res, err := l.client.Get(listsMembersURL, params)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var m ListsMembers
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}

	return &m, nil
}
