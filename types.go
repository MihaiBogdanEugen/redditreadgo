package redditreadgo

// Submission represents an individual post from the perspective of a subreddit
type Submission struct {
	Author        string  `json:"author"`
	Title         string  `json:"title"`
	URL           string  `json:"url"`
	Domain        string  `json:"domain"`
	Subreddit     string  `json:"subreddit"`
	SubredditID   string  `json:"subreddit_id"`
	FullID        string  `json:"name"`
	ID            string  `json:"id"`
	Permalink     string  `json:"permalink"`
	Selftext      string  `json:"selftext"`
	SelftextHTML  string  `json:"selftext_html"`
	ThumbnailURL  string  `json:"thumbnail"`
	DateCreated   float64 `json:"created_utc"`
	NumComments   int     `json:"num_comments"`
	Score         int     `json:"score"`
	Ups           int     `json:"ups"`
	Downs         int     `json:"downs"`
	IsNSFW        bool    `json:"over_18"`
	IsSelf        bool    `json:"is_self"`
	WasClicked    bool    `json:"clicked"`
	IsSaved       bool    `json:"saved"`
	BannedBy      *string `json:"banned_by"`
	LinkFlairText string  `json:"link_flair_text"`
}
