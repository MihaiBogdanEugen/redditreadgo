package redditreadgo

// PopularitySort represents the possible ways to sort submissions by popularity.
type PopularitySort string

const (
	DefaultPopularity        PopularitySort = ""
	HotSubmissions                          = "hot"
	NewSubmissions                          = "new"
	RisingSubmissions                       = "rising"
	TopSubmissions                          = "top"
	ControversialSubmissions                = "controversial"
)

// AgeSort represents the possible ways to sort submissions by age.
type AgeSort string

const (
	ThisHour  AgeSort = "hour"
	ThisDay           = "day"
	ThisWeek          = "week"
	ThisMonth         = "month"
	ThisYear          = "year"
	AllTime           = "all"
)
