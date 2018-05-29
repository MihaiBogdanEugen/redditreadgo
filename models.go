package redditreadgo

// PopularitySort represents the possible ways to sort submissions by popularity.
type PopularitySort string

const (
	// DefaultPopularity
	DefaultPopularity PopularitySort = ""
	// HotSubmissions
	HotSubmissions = "hot"
	// NewSubmissions
	NewSubmissions = "new"
	// RisingSubmissions
	RisingSubmissions = "rising"
	// TopSubmissions
	TopSubmissions = "top"
	// ControversialSubmissions
	ControversialSubmissions = "controversial"
)

// AgeSort represents the possible ways to sort submissions by age.
type AgeSort string

const (
	// ThisHour
	ThisHour AgeSort = "hour"
	// ThisDay
	ThisDay = "day"
	// ThisWeek
	ThisWeek = "week"
	//ThisMonth
	ThisMonth = "month"
	// ThisYear
	ThisYear = "year"
	// AllTime
	AllTime = "all"
)
