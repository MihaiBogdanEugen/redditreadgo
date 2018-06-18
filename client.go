package redditreadgo

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/beefsack/go-rate"
	"github.com/google/go-querystring/query"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

// TokenURL specifies default Reddit access token URL
const TokenURL = "https://www.reddit.com/api/v1/access_token"

// QueryURL specifies default Reddit query URL
const QueryURL = "https://oauth.reddit.com"

const DefaultSliceSize = 100

// ReadOnlyRedditClient represents an OAuth, read-only session with reddit.
type ReadOnlyRedditClient struct {
	Token        *oauth2.Token
	Cookie       *http.Cookie
	clientID     string
	clientSecret string
	userAgent    string
	throttle     *rate.RateLimiter
	logger       *logrus.Logger
}

// IReadOnlyRedditClient defines behaviour for an OAuth, read-only session with reddit.
type IReadOnlyRedditClient interface {

	// Logger sets the logger. Optional, useful for debugging purposes.
	Logger(logger *logrus.Logger)

	// Throttle sets the interval of each HTTP request. Disable by setting interval to 0. Disabled by default.
	Throttle(interval time.Duration)

	// AllSubmissionsTo returns a total no. of submissions to the given subreddit, considering popularity sort and age sort
	AllSubmissionsTo(subreddit string, sort PopularitySort, age AgeSort, total int) ([]*Submission, error)

	// SubmissionsTo returns the submissions to the given subreddit, considering popularity sort, age sort, and listing options
	SubmissionsTo(subreddit string, sort PopularitySort, age AgeSort, params ListingOptions) ([]*Submission, *SliceInfo, error)

	// AllSubmissionsOf returns a total no. of submissions of the given author, considering popularity sort and age sort
	AllSubmissionsOf(author string, sort PopularitySort, age AgeSort, total int) ([]*Submission, error)

	// SubmissionsOf returns the submissions of the given author, considering popularity sort, age sort, and listing options
	SubmissionsOf(author string, sort PopularitySort, age AgeSort, params ListingOptions) ([]*Submission, *SliceInfo, error)
}

// NewReadOnlyRedditClient creates a new session for those who want to log into a reddit account via OAuth.
func NewReadOnlyRedditClient(clientID string, clientSecret string, userAgent string) (*ReadOnlyRedditClient, error) {

	if len(clientID) == 0 {
		return nil, errors.New("clientId must not be null, nor empty")
	}

	if len(clientSecret) == 0 {
		return nil, errors.New("clientSecret must not be null, nor empty")
	}

	if len(userAgent) == 0 {
		return nil, errors.New("userAgent must not be null, nor empty")
	}

	client := &ReadOnlyRedditClient{
		clientID:     clientID,
		clientSecret: clientSecret,
		userAgent:    userAgent,
	}

	if err := client.loginAuth(); err != nil {
		return nil, err
	}

	return client, nil
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

// AllSubmissionsTo returns a total no. of submissions to the given subreddit, considering popularity sort and age sort
func (c *ReadOnlyRedditClient) AllSubmissionsTo(subreddit string, sort PopularitySort, age AgeSort, total int) ([]*Submission, error) {
	return c.getAllSubmissions(subreddit, sort, age, total, c.SubmissionsTo)
}

// SubmissionsTo returns the submissions on the given subreddit, considering popularity sort, age sort, and listing options
func (c *ReadOnlyRedditClient) SubmissionsTo(subreddit string, sort PopularitySort, age AgeSort, params ListingOptions) ([]*Submission, *SliceInfo, error) {

	if len(subreddit) == 0 {
		return nil, nil, errors.New("subreddit cannot be null nor empty")
	}

	queryParams, err := query.Values(params)
	if err != nil {
		return nil, nil, err
	}

	queryParams.Set("t", string(age))
	queryParams.Set("raw_json", strconv.Itoa(1))

	queryURL := fmt.Sprintf("%s/r/%s/%s?%v", QueryURL, subreddit, sort, queryParams.Encode())

	type Response struct {
		Kind string
		Data struct {
			Dist     int
			Children []struct {
				Kind string
				Data *Submission
			}
			After  string
			Before string
		}
	}

	response := new(Response)
	err = c.doGetRequest(queryURL, response)
	if err != nil {
		return nil, nil, err
	}

	submissions := make([]*Submission, len(response.Data.Children))
	for index, child := range response.Data.Children {
		submissions[index] = child.Data
	}

	return submissions, &SliceInfo{Before: response.Data.Before, After: response.Data.After}, nil
}

// AllSubmissionsOf returns a total no. of submissions of the given author, considering popularity sort and age sort
func (c *ReadOnlyRedditClient) AllSubmissionsOf(author string, sort PopularitySort, age AgeSort, total int) ([]*Submission, error) {
	return c.getAllSubmissions(author, sort, age, total, c.SubmissionsOf)
}

// SubmissionsOf returns the submissions on the given author, considering popularity sort, age sort, and listing options
func (c *ReadOnlyRedditClient) SubmissionsOf(author string, sort PopularitySort, age AgeSort, params ListingOptions) ([]*Submission, *SliceInfo, error) {

	if len(author) == 0 {
		return nil, nil, errors.New("author cannot be null nor empty")
	}

	if params.Limit > 100 && c.logger != nil {
		c.logger.Debug("max limit is 100 results - should one need more, `after` or `before` for pagination")
	}

	queryParams, err := query.Values(params)
	if err != nil {
		return nil, nil, err
	}

	if len(sort) > 0 {
		queryParams.Set("sort", string(sort))
	}
	queryParams.Set("t", string(age))
	queryParams.Set("raw_json", strconv.Itoa(1))

	queryURL := fmt.Sprintf("%s/user/%s/submitted?%v", QueryURL, author, queryParams.Encode())

	type Response struct {
		Kind string
		Data struct {
			Dist     int
			Children []struct {
				Kind string
				Data *Submission
			}
			After  string
			Before string
		}
	}

	response := new(Response)
	err = c.doGetRequest(queryURL, response)
	if err != nil {
		return nil, nil, err
	}

	submissions := make([]*Submission, len(response.Data.Children))
	for index, child := range response.Data.Children {
		submissions[index] = child.Data
	}

	return submissions, &SliceInfo{Before: response.Data.Before, After: response.Data.After}, nil
}

func (c *ReadOnlyRedditClient) getAllSubmissions(subredditOrAuthor string, sort PopularitySort, age AgeSort, total int, fn func(string, PopularitySort, AgeSort, ListingOptions) ([]*Submission, *SliceInfo, error)) ([]*Submission, error) {
	if total <= DefaultSliceSize {
		if submissions, _, err := fn(subredditOrAuthor, sort, age, ListingOptions{Limit: total}); err != nil {
			return nil, err
		} else {
			return submissions, nil
		}
	}

	var results []*Submission
	after := ""

	for {
		submissions, slice, err := fn(subredditOrAuthor, sort, age, ListingOptions{
			After: after,
			Limit: DefaultSliceSize,
		})

		if err != nil {
			return nil, err
		}

		for _, submission := range submissions {
			results = append(results, submission)
		}

		if len(results) >= total || len(submissions) == 0 {
			break
		}

		after = slice.After
	}

	return results, nil
}

func (c *ReadOnlyRedditClient) doGetRequest(url string, d interface{}) error {

	if c.logger != nil {
		c.logger.Debugf("doing GET to %s", url)
	}

	if c.throttle != nil {
		if c.logger != nil {
			c.logger.Debugf("must wait")
		}
		c.throttle.Wait()
	}

	if c.Token.Expiry.Before(time.Now().Add(5 * time.Second)) {
		if c.logger != nil {
			c.logger.Debugf("token expired, must fetch a new one")
		}
		if err := c.refreshLoginAuth(); err != nil {
			return err
		}
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	request.Header.Set("Accept", "*/*")
	request.Header.Set("Accept-Encoding", "gzip, deflate")
	request.Header.Set("Authorization", "bearer "+c.Token.AccessToken)
	if c.Cookie != nil {
		request.Header.Set("Cookie", c.Cookie.Name+":"+c.Cookie.Value)
	}
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("User-Agent", c.userAgent)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if code := response.StatusCode; code < 200 || code > 299 {
		return fmt.Errorf("cannot do get request, status: %v", response.Status)
	}

	contentType, _, err := mime.ParseMediaType(response.Header.Get("Content-Type"))
	if err != nil {
		return err
	}

	if contentType != "application/json" {
		return fmt.Errorf("unknown response content type: %s", contentType)
	}

	reader, err := gzip.NewReader(response.Body)
	if err != nil {
		return err
	}
	defer reader.Close()

	responseBody, err := ioutil.ReadAll(io.LimitReader(reader, 1<<20))
	if err != nil {
		return fmt.Errorf("cannot read body of response: %v", err)
	}

	if err = json.Unmarshal(responseBody, d); err != nil {
		return err
	}

	return nil
}

func (c *ReadOnlyRedditClient) loginAuth() error {

	token, cookie, err := c.retrieveTokenAndCookie(url.Values{
		"grant_type": {"client_credentials"},
	})

	if err != nil {
		return err
	}

	c.Token = token
	c.Cookie = cookie

	return nil
}

func (c *ReadOnlyRedditClient) refreshLoginAuth() error {

	if len(c.Token.RefreshToken) == 0 {
		return errors.New("oauth2: token expired and refresh token is not set")
	}

	token, cookie, err := c.retrieveTokenAndCookie(url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {c.Token.RefreshToken},
	})

	if err != nil {
		return err
	}

	c.Token = token
	c.Cookie = cookie

	return nil
}

func (c *ReadOnlyRedditClient) retrieveTokenAndCookie(values url.Values) (*oauth2.Token, *http.Cookie, error) {

	requestBody := strings.NewReader(values.Encode())
	request, err := http.NewRequest("POST", TokenURL, requestBody)
	if err != nil {
		return nil, nil, err
	}

	request.SetBasicAuth(url.QueryEscape(c.clientID), url.QueryEscape(c.clientSecret))
	request.Header.Set("Accept", "*/*")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("User-Agent", c.userAgent)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, nil, err
	}
	defer response.Body.Close()

	if code := response.StatusCode; code < 200 || code > 299 {
		return nil, nil, fmt.Errorf("oauth2: cannot fetch token, status: %v", response.Status)
	}

	contentType, _, err := mime.ParseMediaType(response.Header.Get("Content-Type"))
	if err != nil {
		return nil, nil, err
	}

	if contentType != "application/json" {
		return nil, nil, fmt.Errorf("unknown response content type: %s", contentType)
	}

	responseBody, err := ioutil.ReadAll(io.LimitReader(response.Body, 1<<20))
	if err != nil {
		return nil, nil, fmt.Errorf("oauth2: cannot read body of response: %v", err)
	}

	var tokenAsJSON TokenAsJSON
	if err = json.Unmarshal(responseBody, &tokenAsJSON); err != nil {
		return nil, nil, err
	}

	token := &oauth2.Token{
		AccessToken:  tokenAsJSON.AccessToken,
		TokenType:    tokenAsJSON.TokenType,
		RefreshToken: tokenAsJSON.RefreshToken,
		Expiry:       time.Now().Add(time.Duration(tokenAsJSON.ExpiresIn) * time.Second),
	}

	if len(token.RefreshToken) == 0 {
		token.RefreshToken = values.Get("refresh_token")
	}

	if len(token.AccessToken) == 0 {
		return token, nil, errors.New("oauth2: server response missing access_token")
	}

	if c.logger != nil {
		c.logger.Debugf("got %s access token expiring at %v", token.TokenType, token.Expiry)
	}

	var correctCookie *http.Cookie = nil
	cookies := response.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "edgebucket" {
			correctCookie = cookie
			break
		}
	}
	return token, correctCookie, nil
}
