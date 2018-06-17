package redditreadgo

// PopularitySort represents the possible ways to sort submissions by popularity.
type PopularitySort string

const (
	// DefaultPopularity value
	DefaultPopularity PopularitySort = ""
	// HotSubmissions value
	HotSubmissions = "hot"
	// NewSubmissions value
	NewSubmissions = "new"
	// RisingSubmissions value
	RisingSubmissions = "rising"
	// TopSubmissions value
	TopSubmissions = "top"
	// ControversialSubmissions value
	ControversialSubmissions = "controversial"
)

// AgeSort represents the possible ways to sort submissions by age.
type AgeSort string

const (
	// ThisHour value
	ThisHour AgeSort = "hour"
	// ThisDay value
	ThisDay = "day"
	// ThisWeek value
	ThisWeek = "week"
	//ThisMonth value
	ThisMonth = "month"
	// ThisYear value
	ThisYear = "year"
	// AllTime value
	AllTime = "all"
)
