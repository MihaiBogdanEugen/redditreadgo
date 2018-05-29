package redditreadgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/beefsack/go-rate"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// AccessTokenURL specifies default Reddit access token URL
const AccessTokenURL = "https://www.reddit.com/api/v1/access_token"

// QueryURL specifies default Reddit query URL
const QueryURL = "https://oauth.reddit.com"

// DefaultLimit specifies the default no of retrieved submissions
const DefaultLimit = 100

// ReadOnlyClient represents an OAuth, read-only session with reddit.
type ReadOnlyClient struct {
	Config      *clientcredentials.Config
	HTTPClient  *http.Client
	TokenExpiry time.Time
	throttle    *rate.RateLimiter
	ctx         context.Context
	logger      *logrus.Logger
}

// IReadOnlyClient defines behaviour for an OAuth, read-only session with reddit.
type IReadOnlyClient interface {

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
	SubmissionsTo(subreddit string, sort PopularitySort, age AgeSort, limit int) ([]*Submission, error)

	// Top100SubmissionsOf returns the top submissions on the given author, considering a specified limit
	Top100SubmissionsOf(author string) ([]*Submission, error)

	// SubmissionsOf returns the submissions on the given author, considering popularity sort and a specified limit
	SubmissionsOf(author string, sort PopularitySort, limit int) ([]*Submission, error)

	doGetRequest(link string, d interface{}) error
}

// NewReadOnlyClient creates a new session for those who want to log into a reddit account via OAuth.
func NewReadOnlyClient(clientID string, clientSecret string, userAgent string) *ReadOnlyClient {

	// Inject our custom HTTP client so that a custom headers can be passed during any authentication requests.
	httpClient := &http.Client{
		Transport: &CustomHttpTransport{
			RoundTripper: http.DefaultTransport,
			Headers: map[string]string{
				"Accept":     "*/*",
				"Connection": "keep-alive",
				"User-Agent": fmt.Sprintf("%s PRAW/5.4.0 prawcore/0.14.0", userAgent),
			},
		},
	}

	return &ReadOnlyClient{
		Config: &clientcredentials.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			TokenURL:     AccessTokenURL,
		},
		ctx: context.WithValue(context.Background(), oauth2.HTTPClient, httpClient),
	}
}

// Logger sets the logger. Optional, useful for debugging purposes.
func (c *ReadOnlyClient) Logger(logger *logrus.Logger) {
	c.logger = logger
}

// Throttle sets the interval of each HTTP request. Disable by setting interval to 0. Disabled by default.
func (c *ReadOnlyClient) Throttle(interval time.Duration) {
	if interval == 0 {
		c.throttle = nil
	} else {
		c.throttle = rate.New(1, interval)
	}
}

// LoginAuth creates the a new HTTP client, considering custom headers added, with a new access token.
func (c *ReadOnlyClient) LoginAuth() error {

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
func (c *ReadOnlyClient) Top100SubmissionsAllTimeTo(subreddit string) ([]*Submission, error) {
	return c.SubmissionsTo(subreddit, TopSubmissions, AllTime, 100)
}

// Top100SubmissionsThisYearTo returns top 100 submissions of current year to given subreddit
func (c *ReadOnlyClient) Top100SubmissionsThisYearTo(subreddit string) ([]*Submission, error) {
	return c.SubmissionsTo(subreddit, TopSubmissions, ThisYear, 100)
}

// Top100SubmissionsThisMonthTo returns top 100 submissions of current month to given subreddit
func (c *ReadOnlyClient) Top100SubmissionsThisMonthTo(subreddit string) ([]*Submission, error) {
	return c.SubmissionsTo(subreddit, TopSubmissions, ThisMonth, 100)
}

// Top100SubmissionsThisWeekTo returns top 100 submissions of current week to given subreddit
func (c *ReadOnlyClient) Top100SubmissionsThisWeekTo(subreddit string) ([]*Submission, error) {
	return c.SubmissionsTo(subreddit, TopSubmissions, ThisWeek, 100)
}

// Top100SubmissionsThisDayTo returns top 100 submissions of current day to given subreddit
func (c *ReadOnlyClient) Top100SubmissionsThisDayTo(subreddit string) ([]*Submission, error) {
	return c.SubmissionsTo(subreddit, TopSubmissions, ThisDay, 100)
}

// Top100SubmissionsThisHourTo returns top 100 submissions of current hour to given subreddit
func (c *ReadOnlyClient) Top100SubmissionsThisHourTo(subreddit string) ([]*Submission, error) {
	return c.SubmissionsTo(subreddit, TopSubmissions, ThisHour, 100)
}

// SubmissionsTo returns the submissions on the given subreddit, considering popularity sort, age sort and a specified limit
func (c *ReadOnlyClient) SubmissionsTo(subreddit string, sort PopularitySort, age AgeSort, limit int) ([]*Submission, error) {
	if len(subreddit) == 0 {
		return nil, errors.New("must specify name of subreddit")
	}

	if limit < 1 {
		limit = DefaultLimit
	}

	queryParams := url.Values{}
	queryParams.Set("t", string(age))
	queryParams.Set("limit", strconv.Itoa(limit))

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
	err := c.doGetRequest(queryURL, response)
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
func (c *ReadOnlyClient) Top100SubmissionsOf(author string) ([]*Submission, error) {
	return c.SubmissionsOf(author, TopSubmissions, 100)
}

// SubmissionsOf returns the submissions on the given author, considering popularity sort and a specified limit
func (c *ReadOnlyClient) SubmissionsOf(author string, sort PopularitySort, limit int) ([]*Submission, error) {
	if len(author) == 0 {
		return nil, errors.New("must specify name of the author")
	}

	if limit < 1 {
		limit = DefaultLimit
	}

	queryParams := url.Values{}
	queryParams.Set("sort", string(sort))
	queryParams.Set("limit", strconv.Itoa(limit))
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
	err := c.doGetRequest(queryURL, response)
	if err != nil {
		return nil, err
	}

	submissions := make([]*Submission, len(response.Data.Children))
	for index, child := range response.Data.Children {
		submissions[index] = child.Data
	}

	return submissions, nil
}

func (c *ReadOnlyClient) doGetRequest(url string, d interface{}) error {

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
	defer response.Body.Close()
	if err != nil {
		return err
	}

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