// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crawl

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/ainsleyclark/errors"
	"github.com/krang-backlink/api/common/httputil"
	mocks "github.com/krang-backlink/api/gen/mocks/common/httputil"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestNew(t *testing.T) {
	got := New()
	assert.NotNil(t, got)
}

func ReadTestFile(t *testing.T, file string) []byte {
	t.Helper()

	wd, err := os.Getwd()
	assert.NoError(t, err)

	buf, err := os.ReadFile(filepath.Join(wd, "testdata", file))
	assert.NoError(t, err)

	return buf
}

var (
	// testURL is the default URL used for scrape testing.
	testURL = "https://google.com"
	// defaultContent is the default content when scraped.
	defaultContent = Content{
		H1:          "h1",
		H2:          "h2",
		Title:       "title",
		Description: "description",
		Body:        "body h1 h2",
		SocialImage: testURL + "/image",
	}
)

func TestScrape_Scrape(t *testing.T) {
	tt := map[string]struct {
		url  string
		mock func(client *mocks.Client)
		want any
	}{
		"Bad URL": {
			"@#@#$$%$",
			nil,
			"Bad URL",
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
		"Simple": {
			testURL,
			func(client *mocks.Client) {
				client.On("Do", testURL, http.MethodGet).
					Return(&httputil.Response{Status: http.StatusOK, Body: string(ReadTestFile(t, "simple.html"))}, nil)
			},
			defaultContent,
		},
		"No Quotes": {
			testURL,
			func(client *mocks.Client) {
				client.On("Do", testURL, http.MethodGet).
					Return(&httputil.Response{Status: http.StatusOK, Body: string(ReadTestFile(t, "no-quotes.html"))}, nil)
			},
			defaultContent,
		},
		"Empty": {
			testURL,
			func(client *mocks.Client) {
				client.On("Do", testURL, http.MethodGet).
					Return(&httputil.Response{Status: http.StatusOK, Body: string(ReadTestFile(t, "empty.html"))}, nil)
			},
			Content{},
		},
		"Variations": {
			testURL,
			func(client *mocks.Client) {
				client.On("Do", testURL, http.MethodGet).
					Return(&httputil.Response{Status: http.StatusOK, Body: string(ReadTestFile(t, "variations.html"))}, nil)
			},
			defaultContent,
		},
		"Dirty Content": {
			testURL,
			func(client *mocks.Client) {
				client.On("Do", testURL, http.MethodGet).
					Return(&httputil.Response{Status: http.StatusOK, Body: string(ReadTestFile(t, "dirty-content.html"))}, nil)
			},
			defaultContent,
		},
		"OG Image Variations": {
			testURL,
			func(client *mocks.Client) {
				client.On("Do", testURL, http.MethodGet).
					Return(&httputil.Response{Status: http.StatusOK, Body: string(ReadTestFile(t, "og-image.html"))}, nil)
			},
			Content{Title: "title", Description: "description", Body: "body", SocialImage: "https://krang.com/myimage.jpg"},
		},
		"Relative OG Image": {
			testURL,
			func(client *mocks.Client) {
				client.On("Do", testURL, http.MethodGet).
					Return(&httputil.Response{Status: http.StatusOK, Body: string(ReadTestFile(t, "relative-og-image.html"))}, nil)
			},
			Content{Title: "title", Description: "description", Body: "body", SocialImage: "https://google.com/img/test.jpg"},
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
			got, err := s.Scrape(test.url)
			if err != nil {
				assert.Contains(t, errors.Message(err), test.want)
				return
			}
			assert.Equal(t, test.want, *got)
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
	_, err := s.Scrape(testURL)
	assert.Contains(t, errors.Message(err), "Error creating goquery document")
}
