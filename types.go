package redditreadgo

// PopularitySort represents the possible ways to sort submissions by popularity.
type PopularitySort string

const (
	// DefaultPopularity value
	DefaultPopularity PopularitySort = ""
	// HotSubmissions value
	HotSubmissions PopularitySort = "hot"
	// NewSubmissions value
	NewSubmissions PopularitySort = "new"
	// RisingSubmissions value
	RisingSubmissions PopularitySort = "rising"
	// TopSubmissions value
	TopSubmissions PopularitySort = "top"
	// ControversialSubmissions value
	ControversialSubmissions PopularitySort = "controversial"
)

// AgeSort represents the possible ways to sort submissions by age.
type AgeSort string

const (
	// ThisHour value
	ThisHour AgeSort = "hour"
	// ThisDay value
	ThisDay AgeSort = "day"
	// ThisWeek value
	ThisWeek AgeSort = "week"
	//ThisMonth value
	ThisMonth AgeSort = "month"
	// ThisYear value
	ThisYear AgeSort = "year"
	// AllTime value
	AllTime AgeSort = "all"
)

type Region string

const (
	Global            Region = "GLOBAL"
	USA               Region = "US"
	Argentina         Region = "AR"
	Australia         Region = "AU"
	Bulgaria          Region = "BG"
	Canada            Region = "CA"
	Chile             Region = "CL"
	Colombia          Region = "CO"
	Croatia           Region = "HR"
	CzechRepublic     Region = "CZ"
	Finald            Region = "FI"
	Greece            Region = "GR"
	Hungary           Region = "HU"
	Iceland           Region = "IS"
	India             Region = "IN"
	Ireland           Region = "IE"
	Japan             Region = "JP"
	Malaysia          Region = "MY"
	Mexico            Region = "MX"
	NewZealand        Region = "NZ"
	Philippines       Region = "PH"
	Poland            Region = "PL"
	Portugal          Region = "PT"
	PuertoRico        Region = "PR"
	Romania           Region = "RO"
	Russia            Region = "RS"
	Singapore         Region = "SG"
	Sweden            Region = "SE"
	Taiwan            Region = "TW"
	Thailand          Region = "TH"
	Turkey            Region = "TR"
	UnitedKingdom     Region = "GB"
	USA_Alaska        Region = "US_AK"
	USA_Alabama       Region = "US_AL"
	USA_Arkansas      Region = "US_AR"
	USA_Arizona       Region = "US_AZ"
	USA_California    Region = "US_CA"
	USA_Colorado      Region = "US_CO"
	USA_Connecticut   Region = "US_CT"
	USA_DC            Region = "US_DC"
	USA_Delaware      Region = "US_DE"
	USA_Florida       Region = "US_FL"
	USA_Georgia       Region = "US_GA"
	USA_Hawaii        Region = "US_HI"
	USA_Iowa          Region = "US_IA"
	USA_Idaho         Region = "US_ID"
	USA_Illinois      Region = "US_IL"
	USA_Indiana       Region = "US_IN"
	USA_Kansas        Region = "US_KS"
	USA_Kentucky      Region = "US_KY"
	USA_Louisiana     Region = "US_LA"
	USA_Massachusetts Region = "US_MA"
	USA_Maryland      Region = "US_MD"
	USA_Maine         Region = "US_ME"
	USA_Michigan      Region = "US_MI"
	USA_Minnesota     Region = "US_MN"
	USA_Missouri      Region = "US_MO"
	USA_Mississippi   Region = "US_MS"
	USA_Montana       Region = "US_MT"
	USA_NorthCarolina Region = "US_NC"
	USA_NorthDakota   Region = "US_ND"
	USA_Nebraska      Region = "US_NE"
	USA_NewHampshire  Region = "US_NH"
	USA_NewJersey     Region = "US_NJ"
	USA_NewMexico     Region = "US_NM"
	USA_Nevada        Region = "US_NV"
	USA_NewYork       Region = "US_NY"
	USA_Ohio          Region = "US_OH"
	USA_Oklahoma      Region = "US_OK"
	USA_Oregon        Region = "US_OR"
	USA_Pennsylvania  Region = "US_PA"
	USA_RhodeIsland   Region = "US_RI"
	USA_SouthCarolina Region = "US_SC"
	USA_SouthDakota   Region = "US_SD"
	USA_Tennessee     Region = "US_TN"
	USA_Texas         Region = "US_TX"
	USA_Utah          Region = "US_UT"
	USA_Virginia      Region = "US_VA"
	USA_Vermont       Region = "US_VT"
	USA_Washington    Region = "US_WA"
	USA_Wisconsin     Region = "US_WI"
	USA_WestVirginia  Region = "US_WV"
	USA_Wyoming       Region = "US_WY"
)
