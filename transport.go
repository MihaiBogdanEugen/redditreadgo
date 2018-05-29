package redditreadgo

import "net/http"

// CustomHttpTransport adds the possibility to pass on custom HTTP headers
type CustomHttpTransport struct {
	RoundTripper http.RoundTripper
	Headers      map[string]string
}

// RoundTrip allows for custom headers to be added on any request
func (t *CustomHttpTransport) RoundTrip(request *http.Request) (response *http.Response, err error) {
	for headerName, headerValue := range t.Headers {
		request.Header.Set(headerName, headerValue)
	}
	return t.RoundTripper.RoundTrip(request)
}
