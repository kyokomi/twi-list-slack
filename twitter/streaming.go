package twitter

type Streaming struct {
	Contributors         interface{}      `json:"contributors"`
	Coordinates          interface{}      `json:"coordinates"`
	CreatedAt            string           `json:"created_at"`
	Entities             Entities         `json:"entities"`
	FavoriteCount        int              `json:"favorite_count"`
	Favorited            bool             `json:"favorited"`
	FilterLevel          string           `json:"filter_level"`
	Geo                  interface{}      `json:"geo"`
	ID                   int              `json:"id"`
	IDStr                string           `json:"id_str"`
	InReplyToScreenName  interface{}      `json:"in_reply_to_screen_name"`
	InReplyToStatusID    interface{}      `json:"in_reply_to_status_id"`
	InReplyToStatusIDStr interface{}      `json:"in_reply_to_status_id_str"`
	InReplyToUserID      interface{}      `json:"in_reply_to_user_id"`
	InReplyToUserIDStr   interface{}      `json:"in_reply_to_user_id_str"`
	Lang                 string           `json:"lang"`
	Place                interface{}      `json:"place"`
	RetweetCount         int              `json:"retweet_count"`
	Retweeted            bool             `json:"retweeted"`
	RetweetedStatus      RetweeetedStatus `json:"retweeted_status"`
	Source               string           `json:"source"`
	Text                 string           `json:"text"`
	TimestampMs          string           `json:"timestamp_ms"`
	Truncated            bool             `json:"truncated"`
	User                 User             `json:"user"`
}
