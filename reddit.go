package slackbot

// Reddit is the response data that comes back from the request.
type Reddit struct {
	Kind string    `json:"kind"`
	Data OuterData `json:"data"`
}

// OuterData contains the data for this request query including pagination.
type OuterData struct {
	Modhash  string     `json:"modhash"`
	Children []Children `json:"children"`
	After    string     `json:"after"`
	Before   string     `json:"before"`
}

// Children contains the structure for the inner data.
type Children struct {
	Kind string    `json:"kind"`
	Data InnerData `json:"data"`
}

// InnerData contains the data for each post.
type InnerData struct {
	Likes         int    `json:"likes"`
	ViewCount     int    `json:"view_count"`
	Author        string `json:"author"`
	Score         int    `json:"score"`
	Over18        bool   `json:"over_18"`
	Domain        string `json:"domain"` //youtube.com, other_links, or self.nba
	Downs         int    `json:"downs"`
	URL           string `json:"url"`   //very important!
	Title         string `json:"title"` //also important!
	NumComments   int    `json:"num_comments"`
	SubredditType string `json:"subreddit_type"`
	IsVideo       bool   `json:"is_video"`
	Ups           int    `json:"ups"`
}
