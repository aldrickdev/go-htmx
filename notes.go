package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func makeRequest(url string) (int, string) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	return res.StatusCode, string(body)
}

type loggingRoundTripper struct {
	next   http.RoundTripper
	logger io.Writer
}

func (l loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.logger, "[%s] %s %s\n", time.Now().Format(time.ANSIC), r.Method, r.URL.String())
	return l.next.RoundTrip(r)
}

type retryRoundTripper struct {
	next       http.RoundTripper
	maxRetries int
	delay      time.Duration
}

func (rr retryRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	var attempts int
	for {
		res, err := rr.next.RoundTrip(r)
		attempts++

		// If we have reached our max retry limit return anyways
		if attempts == rr.maxRetries {
			return res, err
		}

		if err == nil && res.StatusCode < http.StatusInternalServerError {
			return res, err
		}

		select {
		case <-r.Context().Done():
			return res, r.Context().Err()

		case <-time.After(rr.delay):
		}
	}
}

type authRoundTripper struct {
	next             http.RoundTripper
	user, pwd, token string
}

func (a authRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	r.SetBasicAuth(a.user, a.pwd)
	token := fmt.Sprintf("bearer %s", a.token)
	r.Header.Set("Authorization", token)
	return a.next.RoundTrip(r)
}

var httpClient = &http.Client{
	// Transport: &loggingRoundTripper{
	// 	next: http.DefaultTransport,
	// 	logger:              os.Stdout,
	// },
	Transport: &authRoundTripper{
		next: &retryRoundTripper{
			next: &loggingRoundTripper{
				next:   http.DefaultTransport,
				logger: os.Stdout,
			},
			maxRetries: 3,
			delay:      time.Duration(1 * time.Second),
		},
		user:  "bobby",
		pwd:   "pwd",
		token: "tokenbaby",
	},
}
