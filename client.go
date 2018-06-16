package redditreadgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/beefsack/go-rate"
	"github.com/google/go-querystring/query"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// AccessTokenURL specifies default Reddit access token URL
const AccessTokenURL = "https://www.reddit.com/api/v1/access_token"

// QueryURL specifies default Reddit query URL
const QueryURL = "https://oauth.reddit.com"

// ListingOptions represents listings query url parameters
// More info: https://www.reddit.com/dev/api/
type ListingOptions struct {

	// Limit - the maximum number of items to return in this slice of the listing.
	Limit int `url:"limit,omitempty"`

	// After or Before - only one should be specified. These indicate the full name of an item in the listing to use as the anchor point of the slice.
	After string `url:"after,omitempty"`

	// Before or After - only one should be specified. These indicate the full name of an item in the listing to use as the anchor point of the slice.
	Before string `url:"before,omitempty"`

	// Count - the number of items already seen in this listing.
	Count int `url:"count,omitempty"`

	// Show - optional parameter; if all is passed, filters such as "hide links that I have voted on" will be disabled.
	Show string `url:"show,omitempty"`
}

// ReadOnlyRedditClient represents an OAuth, read-only session with reddit.
type ReadOnlyRedditClient struct {
	Config      *clientcredentials.Config
	HTTPClient  *http.Client
	TokenExpiry time.Time
	throttle    *rate.RateLimiter
	ctx         context.Context
	logger      *logrus.Logger
}

// IReadOnlyRedditClient defines behaviour for an OAuth, read-only session with reddit.
type IReadOnlyRedditClient interface {

	// Logger sets the logger. Optional, useful for debugging purposes.
	Logger(logger *logrus.Logger)

	// Throttle sets the interval of each HTTP request. Disable by setting interval to 0. Disabled by default.
	Throttle(interval time.Duration)

	// LoginAuth creates the a new HTTP client, considering custom headers added, with a new access token.
	LoginAuth() error

	// Top100SubmissionsAllTimeTo returns top 100 submissions of all time to given subreddit
	Top100SubmissionsAllTimeTo(subreddit string) ([]*Submission, error)

	// Top100SubmissionsAllTimeTo returns top 100 submissions of current year to given subreddit
	Top100SubmissionsThisYearTo(subreddit string) ([]*Submission, error)

	// Top100SubmissionsThisMonthTo returns top 100 submissions of current month to given subreddit
	Top100SubmissionsThisMonthTo(subreddit string) ([]*Submission, error)

	// Top100SubmissionsThisWeekTo returns top 100 submissions of current week to given subreddit
	Top100SubmissionsThisWeekTo(subreddit string) ([]*Submission, error)

	// Top100SubmissionsThisDayTo returns top 100 submissions of current day to given subreddit
	Top100SubmissionsThisDayTo(subreddit string) ([]*Submission, error)

	// Top100SubmissionsThisHourTo returns top 100 submissions of current hour to given subreddit
	Top100SubmissionsThisHourTo(subreddit string) ([]*Submission, error)

	// SubmissionsTo returns the submissions to the given subreddit, considering popularity sort, age sort and a specified limit
	SubmissionsTo(subreddit string, sort PopularitySort, age AgeSort, params ListingOptions) ([]*Submission, error)

	// Top100SubmissionsOf returns the top submissions on the given author, considering a specified limit
	Top100SubmissionsOf(author string) ([]*Submission, error)

	// SubmissionsOf returns the submissions on the given author, considering popularity sort and a specified limit
	SubmissionsOf(author string, sort PopularitySort, params ListingOptions) ([]*Submission, error)

	doGetRequest(link string, d interface{}) error
}

// NewReadOnlyRedditClient creates a new session for those who want to log into a reddit account via OAuth.
func NewReadOnlyRedditClient(clientID string, clientSecret string, userAgent string) *ReadOnlyRedditClient {

	// Inject our custom HTTP client so that a custom headers can be passed during any authentication requests.
	httpClient := &http.Client{
		Transport: &CustomHttpTransport{
			RoundTripper: http.DefaultTransport,
			Headers: map[string]string{
				"Accept":     "*/*",
				"Connection": "keep-alive",
				"User-Agent": userAgent,
			},
		},
	}

	return &ReadOnlyRedditClient{
		Config: &clientcredentials.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			TokenURL:     AccessTokenURL,
		},
		ctx: context.WithValue(context.Background(), oauth2.HTTPClient, httpClient),
	}
}

// Logger sets the logger. Optional, useful for debugging purposes.
func (c *ReadOnlyRedditClient) Logger(logger *logrus.Logger) {
	c.logger = logger
}

// Throttle sets the interval of each HTTP request. Disable by setting interval to 0. Disabled by default.
func (c *ReadOnlyRedditClient) Throttle(interval time.Duration) {
	if interval == 0 {
		c.throttle = nil
	} else {
		c.throttle = rate.New(1, interval)
	}
}

// LoginAuth creates the a new HTTP client, considering custom headers added, with a new access token.
func (c *ReadOnlyRedditClient) LoginAuth() error {

	token, err := c.Config.Token(c.ctx)
	if err != nil {
		return err
	}
	if !token.Valid() {
		msg := "Invalid OAuth token"
		if token != nil {
			if extra := token.Extra("error"); extra != nil {
				msg = fmt.Sprintf("%s: %s", msg, extra)
			}
		}
		return errors.New(msg)
	}

	if c.logger != nil {
		c.logger.Debugf("got %s access token expiring at %v", token.TokenType, token.Expiry)
	}

	c.TokenExpiry = token.Expiry
	c.HTTPClient = c.Config.Client(c.ctx)

	return nil
}

// Top100SubmissionsAllTimeTo returns top 100 submissions of all time to given subreddit
func (c *ReadOnlyRedditClient) Top100SubmissionsAllTimeTo(subreddit string) ([]*Submission, error) {
	return c.SubmissionsTo(subreddit, TopSubmissions, AllTime, ListingOptions{Limit: 100})
}

// Top100SubmissionsThisYearTo returns top 100 submissions of current year to given subreddit
func (c *ReadOnlyRedditClient) Top100SubmissionsThisYearTo(subreddit string) ([]*Submission, error) {
	return c.SubmissionsTo(subreddit, TopSubmissions, ThisYear, ListingOptions{Limit: 100})
}

// Top100SubmissionsThisMonthTo returns top 100 submissions of current month to given subreddit
func (c *ReadOnlyRedditClient) Top100SubmissionsThisMonthTo(subreddit string) ([]*Submission, error) {
	return c.SubmissionsTo(subreddit, TopSubmissions, ThisMonth, ListingOptions{Limit: 100})
}

// Top100SubmissionsThisWeekTo returns top 100 submissions of current week to given subreddit
func (c *ReadOnlyRedditClient) Top100SubmissionsThisWeekTo(subreddit string) ([]*Submission, error) {
	return c.SubmissionsTo(subreddit, TopSubmissions, ThisWeek, ListingOptions{Limit: 100})
}

// Top100SubmissionsThisDayTo returns top 100 submissions of current day to given subreddit
func (c *ReadOnlyRedditClient) Top100SubmissionsThisDayTo(subreddit string) ([]*Submission, error) {
	return c.SubmissionsTo(subreddit, TopSubmissions, ThisDay, ListingOptions{Limit: 100})
}

// Top100SubmissionsThisHourTo returns top 100 submissions of current hour to given subreddit
func (c *ReadOnlyRedditClient) Top100SubmissionsThisHourTo(subreddit string) ([]*Submission, error) {
	return c.SubmissionsTo(subreddit, TopSubmissions, ThisHour, ListingOptions{Limit: 100})
}

// SubmissionsTo returns the submissions on the given subreddit, considering popularity sort, age sort and a specified limit
func (c *ReadOnlyRedditClient) SubmissionsTo(subreddit string, sort PopularitySort, age AgeSort, params ListingOptions) ([]*Submission, error) {
	if len(subreddit) == 0 {
		return nil, errors.New("must specify name of subreddit")
	}

	queryParams, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	queryParams.Set("t", string(age))

	queryURL := fmt.Sprintf("%s/r/%s/%s.json?%v", QueryURL, subreddit, sort, queryParams.Encode())

	if c.logger != nil {
		c.logger.Debugf("queryURL = %s", queryURL)
	}

	type Response struct {
		Data struct {
			Children []struct {
				Data *Submission
			}
		}
	}

	response := new(Response)
	err = c.doGetRequest(queryURL, response)
	if err != nil {
		return nil, err
	}

	submissions := make([]*Submission, len(response.Data.Children))
	for index, child := range response.Data.Children {
		submissions[index] = child.Data
	}

	return submissions, nil
}

// Top100SubmissionsOf returns the top submissions on the given author, considering a specified limit
func (c *ReadOnlyRedditClient) Top100SubmissionsOf(author string) ([]*Submission, error) {
	return c.SubmissionsOf(author, TopSubmissions, ListingOptions{Limit: 100})
}

// SubmissionsOf returns the submissions on the given author, considering popularity sort and a specified limit
func (c *ReadOnlyRedditClient) SubmissionsOf(author string, sort PopularitySort, params ListingOptions) ([]*Submission, error) {
	if len(author) == 0 {
		return nil, errors.New("must specify name of the author")
	}

	queryParams, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	if len(sort) > 0 {
		queryParams.Set("sort", string(sort))
	}

	queryParams.Set("raw_json", strconv.Itoa(1))

	queryURL := fmt.Sprintf("%s/user/%s/submitted?%v", QueryURL, author, queryParams.Encode())

	if c.logger != nil {
		c.logger.Debugf("queryURL = %s", queryURL)
	}

	type Response struct {
		Data struct {
			Children []struct {
				Data *Submission
			}
		}
	}

	response := new(Response)
	err = c.doGetRequest(queryURL, response)
	if err != nil {
		return nil, err
	}

	submissions := make([]*Submission, len(response.Data.Children))
	for index, child := range response.Data.Children {
		submissions[index] = child.Data
	}

	return submissions, nil
}

func (c *ReadOnlyRedditClient) doGetRequest(url string, d interface{}) error {

	if c.HTTPClient == nil {
		return errors.New("no HttpClient - use LoginAuth to create one")
	}

	if c.throttle != nil {
		c.throttle.Wait()
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, d)
	if err != nil {
		return err
	}

	return nil
}
