package twitter

// User twitter account user
type User struct {
	ContributorsEnabled            bool        `json:"contributors_enabled"`
	CreatedAt                      string      `json:"created_at"`
	DefaultProfile                 bool        `json:"default_profile"`
	DefaultProfileImage            bool        `json:"default_profile_image"`
	Description                    string      `json:"description"`
	Entities                       Entities    `json:"entities"`
	FavouritesCount                int         `json:"favourites_count"`
	FollowRequestSent              bool        `json:"follow_request_sent"`
	FollowersCount                 int         `json:"followers_count"`
	Following                      bool        `json:"following"`
	FriendsCount                   int         `json:"friends_count"`
	GeoEnabled                     bool        `json:"geo_enabled"`
	ID                             int         `json:"id"`
	IDStr                          string      `json:"id_str"`
	IsTranslationEnabled           bool        `json:"is_translation_enabled"`
	IsTranslator                   bool        `json:"is_translator"`
	Lang                           string      `json:"lang"`
	ListedCount                    int         `json:"listed_count"`
	Location                       string      `json:"location"`
	Name                           string      `json:"name"`
	Notifications                  bool        `json:"notifications"`
	ProfileBackgroundColor         string      `json:"profile_background_color"`
	ProfileBackgroundImageURL      string      `json:"profile_background_image_url"`
	ProfileBackgroundImageURLHTTPS string      `json:"profile_background_image_url_https"`
	ProfileBackgroundTile          bool        `json:"profile_background_tile"`
	ProfileBannerURL               string      `json:"profile_banner_url"`
	ProfileImageURL                string      `json:"profile_image_url"`
	ProfileImageURLHTTPS           string      `json:"profile_image_url_https"`
	ProfileLinkColor               string      `json:"profile_link_color"`
	ProfileLocation                interface{} `json:"profile_location"`
	ProfileSidebarBorderColor      string      `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor        string      `json:"profile_sidebar_fill_color"`
	ProfileTextColor               string      `json:"profile_text_color"`
	ProfileUseBackgroundImage      bool        `json:"profile_use_background_image"`
	Protected                      bool        `json:"protected"`
	ScreenName                     string      `json:"screen_name"`
	Status                         Status      `json:"status"`
	StatusesCount                  int         `json:"statuses_count"`
	TimeZone                       string      `json:"time_zone"`
	URL                            string      `json:"url"`
	UtcOffset                      int         `json:"utc_offset"`
	Verified                       bool        `json:"verified"`
}
