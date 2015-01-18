package twitter

type Entities struct {
	Hashtags    []interface{} `json:"hashtags"`
	Medias      []Media       `json:"media"`
	Symbols     []interface{} `json:"symbols"`
	Description struct {
		Urls []URLs `json:"urls"`
	} `json:"description"`
	URL struct {
		Urls []URLs `json:"urls"`
	} `json:"url"`
	UserMentions []interface{} `json:"user_mentions"`
}

type URLs struct {
	DisplayURL  string `json:"display_url"`
	ExpandedURL string `json:"expanded_url"`
	Indices     []int  `json:"indices"`
	URL         string `json:"url"`
}

type Media struct {
	DisplayURL    string `json:"display_url"`
	ExpandedURL   string `json:"expanded_url"`
	ID            int    `json:"id"`
	IDStr         string `json:"id_str"`
	Indices       []int  `json:"indices"`
	MediaURL      string `json:"media_url"`
	MediaURLHTTPS string `json:"media_url_https"`
	Sizes         struct {
		Large struct {
			H      int    `json:"h"`
			Resize string `json:"resize"`
			W      int    `json:"w"`
		} `json:"large"`
		Medium struct {
			H      int    `json:"h"`
			Resize string `json:"resize"`
			W      int    `json:"w"`
		} `json:"medium"`
		Small struct {
			H      int    `json:"h"`
			Resize string `json:"resize"`
			W      int    `json:"w"`
		} `json:"small"`
		Thumb struct {
			H      int    `json:"h"`
			Resize string `json:"resize"`
			W      int    `json:"w"`
		} `json:"thumb"`
	} `json:"sizes"`
	Type string `json:"type"`
	URL  string `json:"url"`
}
