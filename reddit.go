package main

type Reddit struct {
	Kind string `json:"kind"`
	Data struct {
		Modhash  string `json:"modhash"`
		Children []struct {
			Kind string `json:"kind"`
			Data struct {
				ContestMode   bool        `json:"contest_mode"`
				ApprovedAtUtc interface{} `json:"approved_at_utc"`
				BannedBy      interface{} `json:"banned_by"`
				MediaEmbed    struct {
				} `json:"media_embed"`
				Subreddit        string        `json:"subreddit"`
				SelftextHTML     string        `json:"selftext_html"`
				Selftext         string        `json:"selftext"`
				Likes            interface{}   `json:"likes"`
				SuggestedSort    interface{}   `json:"suggested_sort"`
				UserReports      []interface{} `json:"user_reports"`
				SecureMedia      interface{}   `json:"secure_media"`
				LinkFlairText    interface{}   `json:"link_flair_text"`
				ID               string        `json:"id"`
				BannedAtUtc      interface{}   `json:"banned_at_utc"`
				ViewCount        interface{}   `json:"view_count"`
				SecureMediaEmbed struct {
				} `json:"secure_media_embed"`
				Clicked               bool          `json:"clicked"`
				ReportReasons         interface{}   `json:"report_reasons"`
				Author                string        `json:"author"`
				Saved                 bool          `json:"saved"`
				ModReports            []interface{} `json:"mod_reports"`
				CanModPost            bool          `json:"can_mod_post"`
				Name                  string        `json:"name"`
				Score                 int           `json:"score"`
				ApprovedBy            interface{}   `json:"approved_by"`
				Over18                bool          `json:"over_18"`
				Domain                string        `json:"domain"`
				Hidden                bool          `json:"hidden"`
				Thumbnail             string        `json:"thumbnail"`
				SubredditID           string        `json:"subreddit_id"`
				Edited                int           `json:"edited"`
				LinkFlairCSSClass     interface{}   `json:"link_flair_css_class"`
				AuthorFlairCSSClass   string        `json:"author_flair_css_class"`
				Gilded                int           `json:"gilded"`
				Downs                 int           `json:"downs"`
				BrandSafe             bool          `json:"brand_safe"`
				Archived              bool          `json:"archived"`
				RemovalReason         interface{}   `json:"removal_reason"`
				CanGild               bool          `json:"can_gild"`
				IsSelf                bool          `json:"is_self"`
				HideScore             bool          `json:"hide_score"`
				Spoiler               bool          `json:"spoiler"`
				Permalink             string        `json:"permalink"`
				NumReports            interface{}   `json:"num_reports"`
				Locked                bool          `json:"locked"`
				Stickied              bool          `json:"stickied"`
				Created               float64       `json:"created"`
				URL                   string        `json:"url"`
				AuthorFlairText       string        `json:"author_flair_text"`
				Quarantine            bool          `json:"quarantine"`
				Title                 string        `json:"title"`
				CreatedUtc            float64       `json:"created_utc"`
				SubredditNamePrefixed string        `json:"subreddit_name_prefixed"`
				Distinguished         interface{}   `json:"distinguished"`
				Media                 interface{}   `json:"media"`
				NumComments           int           `json:"num_comments"`
				Visited               bool          `json:"visited"`
				SubredditType         string        `json:"subreddit_type"`
				IsVideo               bool          `json:"is_video"`
				Ups                   int           `json:"ups"`
			} `json:"data"`
		} `json:"children"`
		After  string      `json:"after"`
		Before interface{} `json:"before"`
	} `json:"data"`
}
