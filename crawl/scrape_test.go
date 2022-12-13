// Copyright 2023 Ainsley Clark. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package crawl

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/ainsleyclark/errors"
	"github.com/ainsleyclark/stock-informer/httputil"
	mocks "github.com/ainsleyclark/stock-informer/mocks/httputil"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	got := New()
	assert.NotNil(t, got)
}

var (
	// testURL is the default URL used for scrape testing.
	testURL = "https://google.com"
	// testElement is the default selector used for scrape testing.
	testElement = ".element"
)

func TestScrape_Scrape(t *testing.T) {
	tt := map[string]struct {
		url  string
		mock func(client *mocks.Client)
		want any
	}{
		"OK": {
			testURL,
			func(client *mocks.Client) {
				client.On("Do", testURL, http.MethodGet).
					Return(&httputil.Response{
						Status: http.StatusOK,
						Body:   `<html><div class="element">Contents</div></html>`,
					}, nil)
			},
			"Contents",
		},
		"Client Error": {
			testURL,
			func(client *mocks.Client) {
				client.On("Do", testURL, http.MethodGet).
					Return(nil, &errors.Error{Message: "do error"})
			},
			"do error",
		},
		"Bad Status Code": {
			testURL,
			func(client *mocks.Client) {
				client.On("Do", testURL, http.MethodGet).
					Return(&httputil.Response{Status: http.StatusBadRequest}, nil)
			},
			"Error scraping page, status code: 400",
		},
		"No Element": {
			testURL,
			func(client *mocks.Client) {
				client.On("Do", testURL, http.MethodGet).
					Return(&httputil.Response{
						Status: http.StatusOK,
						Body:   "<html></html>",
					}, nil)
			},
			"Error finding element with selector",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			m := &mocks.Client{}
			if test.mock != nil {
				test.mock(m)
			}
			s := scrape{
				client:      m,
				newDocument: goquery.NewDocumentFromReader,
			}
			got, err := s.Scrape(test.url, testElement)
			if err != nil {
				assert.Contains(t, errors.Message(err), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}

func TestScrape_Scrape_ErrorDocument(t *testing.T) {
	mf := func(client *mocks.Client) {
		client.On("Do", testURL, http.MethodGet).
			Return(&httputil.Response{Status: http.StatusOK, Body: "hello"}, nil)
	}
	m := &mocks.Client{}
	mf(m)
	s := scrape{
		client: m,
		newDocument: func(r io.Reader) (*goquery.Document, error) {
			return nil, errors.New("error")
		},
	}
	_, err := s.Scrape(testURL, testElement)
	assert.Contains(t, errors.Message(err), "Error reading document")
}
