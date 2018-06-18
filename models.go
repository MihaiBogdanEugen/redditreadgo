package redditreadgo

// Submission represents an individual post from the perspective of a subreddit
type Submission struct {
	ApprovedAtUTC         float64 `json:"approved_at_utc"`
	ApprovedBy            string  `json:"approved_by"`
	Archived              bool    `json:"archived"`
	Author                string  `json:"author"`
	BannedAtUTC           float64 `json:"banned_at_utc"`
	BannedBy              string  `json:"banned_by"`
	CanGlid               bool    `json:"can_gild"`
	Category              string  `json:"category"`
	Clicked               bool    `json:"clicked"`
	ContentCategories     string  `json:"content_categories"`
	ContestMode           bool    `json:"contest_mode"`
	Created               float64 `json:"created"`
	CreatedUTC            float64 `json:"created_utc"`
	Distinguished         string  `json:"distinguished"`
	Domain                string  `json:"domain"`
	Downs                 int     `json:"downs"`
	Edited                bool    `json:"edited"`
	Glided                uint64  `json:"gilded"`
	Hidden                bool    `json:"hidden"`
	HideScore             bool    `json:"hide_score"`
	ID                    string  `json:"id"`
	IsCrosspostable       bool    `json:"is_crosspostable"`
	IsOriginalContent     bool    `json:"is_original_content"`
	IsRedditMediaDomain   bool    `json:"is_reddit_media_domain"`
	IsSelf                bool    `json:"is_self"`
	IsVideo               bool    `json:"is_video"`
	Likes                 string  `json:"likes"`
	Locked                bool    `json:"locked"`
	MediaOnly             bool    `json:"media_only"`
	Name                  string  `json:"name"`
	NoFollow              bool    `json:"no_follow"`
	NumComments           uint64  `json:"num_comments"`
	NumCrossposts         uint64  `json:"num_crossposts"`
	NumReports            uint64  `json:"num_reports"`
	Over18                bool    `json:"over_18"`
	ParentWhitelistStatus string  `json:"parent_whitelist_status"`
	Permalink             string  `json:"permalink"`
	Pinned                bool    `json:"pinned"`
	PostCategories        string  `json:"post_categories"`
	PostHint              string  `json:"post_hint"`
	Quarantine            bool    `json:"quarantine"`
	RemovalReason         string  `json:"removal_reason"`
	ReportReasons         string  `json:"report_reasons"`
	Saved                 bool    `json:"saved"`
	Score                 uint64  `json:"score"`
	Selftext              string  `json:"selftext"`
	SelftextHTML          string  `json:"selftext_html"`
	SendReplies           bool    `json:"send_replies"`
	Spoiler               bool    `json:"spoiler"`
	Stickied              bool    `json:"stickied"`
	Subreddit             string  `json:"subreddit"`
	SubredditID           string  `json:"subreddit_id"`
	SubredditNamePrefixed string  `json:"subreddit_name_prefixed"`
	SubredditSubscribers  uint64  `json:"subreddit_subscribers"`
	SubredditType         string  `json:"subreddit_type"`
	SuggestedSort         string  `json:"suggested_sort"`
	Thumbnail             string  `json:"thumbnail"`
	Title                 string  `json:"title"`
	Ups                   int     `json:"ups"`
	URL                   string  `json:"url"`
	ViewCount             uint64  `json:"view_count"`
	Visited               bool    `json:"visited"`
	WhitelistStatus       string  `json:"whitelist_status"`
}

// TokenAsJSON represents the access token serialized as a json object
type TokenAsJSON struct {
	// AccessToken value
	AccessToken string `json:"access_token"`
	// TokenType value
	TokenType string `json:"token_type"`
	// RefreshToken value
	RefreshToken string `json:"refresh_token"`
	// ExpiresIn value
	ExpiresIn int32 `json:"expires_in"`
}

type SliceInfo struct {
	After  string
	Before string
}

// ListingOptions represents listings query url parameters. More info: https://www.reddit.com/dev/api/
type ListingOptions struct {
	// Region - filter hot results by specifying the region
	Region Region `url:"q,omitempty"`

	// Limit - the maximum number of items to return in this slice of the listing - default: 25, maximum: 100
	Limit int `url:"limit,omitempty"`

	// After or Before - only one should be specified. These indicate the full name of an item in the listing to use as the anchor point of the slice
	After string `url:"after,omitempty"`

	// Before or After - only one should be specified. These indicate the full name of an item in the listing to use as the anchor point of the slice
	Before string `url:"before,omitempty"`

	// Count - the number of items already seen in this listing - default: 0
	Count int `url:"count,omitempty"`

	// Show - optional parameter; if all is passed, filters such as "hide links that I have voted on" will be disabled
	Show string `url:"show,omitempty"`
}
