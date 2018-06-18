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

// Region represents the possible values for querying by region
type Region string

const (
	Global           Region = "GLOBAL"
	USA              Region = "US"
	Argentina        Region = "AR"
	Australia        Region = "AU"
	Bulgaria         Region = "BG"
	Canada           Region = "CA"
	Chile            Region = "CL"
	Colombia         Region = "CO"
	Croatia          Region = "HR"
	CzechRepublic    Region = "CZ"
	Finald           Region = "FI"
	Greece           Region = "GR"
	Hungary          Region = "HU"
	Iceland          Region = "IS"
	India            Region = "IN"
	Ireland          Region = "IE"
	Japan            Region = "JP"
	Malaysia         Region = "MY"
	Mexico           Region = "MX"
	NewZealand       Region = "NZ"
	Philippines      Region = "PH"
	Poland           Region = "PL"
	Portugal         Region = "PT"
	PuertoRico       Region = "PR"
	Romania          Region = "RO"
	Russia           Region = "RS"
	Singapore        Region = "SG"
	Sweden           Region = "SE"
	Taiwan           Region = "TW"
	Thailand         Region = "TH"
	Turkey           Region = "TR"
	UnitedKingdom    Region = "GB"
	USAAlaska        Region = "US_AK"
	USAAlabama       Region = "US_AL"
	USAArkansas      Region = "US_AR"
	USAArizona       Region = "US_AZ"
	USACalifornia    Region = "US_CA"
	USAColorado      Region = "US_CO"
	USAConnecticut   Region = "US_CT"
	USADC            Region = "US_DC"
	USADelaware      Region = "US_DE"
	USAFlorida       Region = "US_FL"
	USAGeorgia       Region = "US_GA"
	USAHawaii        Region = "US_HI"
	USAIowa          Region = "US_IA"
	USAIdaho         Region = "US_ID"
	USAIllinois      Region = "US_IL"
	USAIndiana       Region = "US_IN"
	USAKansas        Region = "US_KS"
	USAKentucky      Region = "US_KY"
	USALouisiana     Region = "US_LA"
	USAMassachusetts Region = "US_MA"
	USAMaryland      Region = "US_MD"
	USAMaine         Region = "US_ME"
	USAMichigan      Region = "US_MI"
	USAMinnesota     Region = "US_MN"
	USAMissouri      Region = "US_MO"
	USAMississippi   Region = "US_MS"
	USAMontana       Region = "US_MT"
	USANorthCarolina Region = "US_NC"
	USANorthDakota   Region = "US_ND"
	USANebraska      Region = "US_NE"
	USANewHampshire  Region = "US_NH"
	USANewJersey     Region = "US_NJ"
	USANewMexico     Region = "US_NM"
	USANevada        Region = "US_NV"
	USANewYork       Region = "US_NY"
	USAOhio          Region = "US_OH"
	USAOklahoma      Region = "US_OK"
	USAOregon        Region = "US_OR"
	USAPennsylvania  Region = "US_PA"
	USARhodeIsland   Region = "US_RI"
	USASouthCarolina Region = "US_SC"
	USASouthDakota   Region = "US_SD"
	USATennessee     Region = "US_TN"
	USATexas         Region = "US_TX"
	USAUtah          Region = "US_UT"
	USAVirginia      Region = "US_VA"
	USAVermont       Region = "US_VT"
	USAWashington    Region = "US_WA"
	USAWisconsin     Region = "US_WI"
	USAWestVirginia  Region = "US_WV"
	USAWyoming       Region = "US_WY"
)
