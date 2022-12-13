// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package httputil

import (
	"github.com/Danny-Dasilva/CycleTLS/cycletls"
	"github.com/ainsleyclark/errors"
)

type (
	// Client defines the single method interface for creating
	// and performing requests.
	Client interface {
		// Do creates and performs single http Request using the
		// application default client.
		//
		// Returns errors.INTERNAL if the request could not be
		// performed.
		Do(URL, method string) (*Response, error)
	}
	// cycleTLS is an abstraction of the Do method
	// to perform http Requests.
	cycleTLS interface {
		Do(string, cycletls.Options, string) (cycletls.Response, error)
	}
	//Response contains the client response data from
	// the request.
	Response struct {
		RequestID string
		Status    int
		Body      string
		Headers   map[string]string
		Location  string
	}
	// httpClient represents the cycleTLS client used
	// for performing requests.
	httpClient struct {
		cycle cycleTLS
	}
)

const (
	// ClientTimeout specifies a time limit for requests
	// made by the Client in seconds.
	ClientTimeout = 30
	// UserAgent is the user agent header used for making
	// http requests.
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36"
)

var (
	// options defines the cycleTLS options used for creating
	// http requests.
	options = cycletls.Options{
		Headers: map[string]string{
			"User-Agent":      UserAgent,
			"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
			"Accept-Encoding": "gzip, deflate",
			"Accept-Language": "en-GB,en-US;q=0.9,en;q=0.8",
			"Cache-Control":   "no-cache",
		},
		Ja3:             "771,4865-4867-4866-49195-49199-52393-52392-49196-49200-49162-49161-49171-49172-51-57-47-53-10,0-23-65281-10-11-35-16-5-51-43-13-45-28-21,29-23-24-25-256-257,0",
		UserAgent:       UserAgent,
		Cookies:         nil,
		Timeout:         ClientTimeout,
		DisableRedirect: true,
		HeaderOrder: []string{
			"User-Agent",
			"Accept",
			"Accept-Encoding",
			"Accept-Language",
			"Cache-Control",
		},
	}
)

// NewClient returns the common Client used across the
// application
func NewClient() Client {
	return &httpClient{
		cycle: cycletls.Init(),
	}
}

// Do creates and performs single http Request using the
// application default client.
//
// Returns errors.INTERNAL if the request could not be
// performed.
func (h *httpClient) Do(url, method string) (*Response, error) {
	const op = "HTTPClient.Do"
	resp, err := h.cycle.Do(url, options, method)

	// Setup the response data.
	response := &Response{
		RequestID: resp.RequestID,
		Status:    resp.Status,
		Body:      resp.Body,
		Headers:   resp.Headers,
		Location:  resp.Headers["Location"],
	}

	// Handle redirects.
	if Is3xx(resp.Status) && response.Location != "" {
		location, err := GetAbsoluteURL(url, response.Location)
		if err != nil {
			return response, err
		}
		// Bail, the URL and the location header is the same.
		// It could cause an infinite loop
		if location == url {
			return response, nil
		}

		// Recursively call this function with the new
		// location header.
		return h.Do(location, method)
	}

	// Bail if the error is not nil
	if err != nil {
		return nil, errors.NewInternal(err, "Error performing client request", op)
	}

	return response, nil
}
