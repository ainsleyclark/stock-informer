// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crawl

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/ainsleyclark/errors"
	"github.com/ainsleyclark/stock-informer/httputil"
	"io"
	"net/http"
	"net/url"
)

type (
	// Scraper defines the method used for crawling a
	// singular URL.
	Scraper interface {
		// Scrape crawls a webpage obtaining content from the
		// website by the given URL.
		// A http.MethodGet request is made to the URL and the
		// page is analysed for content providing a 200 response
		// is returned.
		//
		// Returns errors.INVALID if the URL failed to parse, the client
		// could not make request or if the status code is not 200.
		// Returns errors.INTERNAL if the document could not be parsed
		// body could not be read.
		Scrape(url string) (*Content, error)
	}
	// Content represents the data returned by the
	// scrape of a URL.
	Content struct {
		H1          string
		H2          string
		Title       string
		Description string
		Body        string
		SocialImage string
	}
	// scrape implements the scraper interface for crawling URLs.
	scrape struct {
		client      httputil.Client
		newDocument func(r io.Reader) (*goquery.Document, error)
	}
)

// New creates a new scraper with a custom http.Client and
// cookie Jar.
// Returns errors.INTERNAL if the jar could not be created.
func New() Scraper {
	return &scrape{
		client:      httputil.NewClient(),
		newDocument: goquery.NewDocumentFromReader,
	}
}

func (s *scrape) Scrape(uri string) (*Content, error) {
	const op = "Scraper.Scrape"

	// Parse the URL
	uriParsed, err := url.Parse(uri)
	if err != nil {
		return nil, errors.NewInvalid(err, "Bad URL", op)
	}

	response, err := s.client.Do(uri, http.MethodGet)
	if err != nil {
		return nil, err
	}

	if !httputil.Is2xx(response.Status) {
		return nil, errors.NewInvalid(errors.New("bad status code"), fmt.Sprintf("Error scraping page, status code: %d", response.Status), op)
	}

	// Load the HTML document from the reader.
	d, err := s.newDocument(bytes.NewBuffer([]byte(response.Body)))
	if err != nil {
		return nil, errors.NewInternal(err, "Error creating goquery document", op)
	}

	doc := document{Doc: d, URL: uriParsed}
	doc.Sanitise()
	return &Content{
		H1:          doc.H1(),
		H2:          doc.H2(),
		Title:       doc.Title(),
		Description: doc.Description(),
		Body:        doc.Body(),
		SocialImage: doc.SocialImage(),
	}, nil
}
