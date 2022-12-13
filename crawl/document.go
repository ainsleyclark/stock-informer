// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crawl

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/krang-backlink/api/common/httputil"
	"github.com/krang-backlink/api/common/stringutil"
	"net/url"
	"strings"
)

// document represents a goquery document and implements
// helper methods to retrieve key information from the
// project page.
type document struct {
	Doc *goquery.Document
	URL *url.URL
}

// Sanitise strips the scrape of unwanted data.
func (d *document) Sanitise() {
	d.Doc.Find("script").Remove()
	d.Doc.Find("style").Remove()
	d.Doc.Find("noscript").Remove()
	d.Doc.Find("form").Remove()
	d.Doc.Find("button").Remove()
	d.Doc.Find("select").Remove()
	d.Doc.Find("iframe").Remove()
}

// Title returns the meta <title> from the document.
func (d *document) Title() string {
	return sanitize(d.Doc.Find("title").First().Text())
}

// Description returns the meta description from the document.
func (d *document) Description() string {
	variations := []string{
		`meta[name=description]`,
		`meta[name=Description]`,
		`meta[name="description"]`,
		`meta[name="Description"]`,
	}
	for _, variation := range variations {
		title, ok := d.Doc.Find(variation).First().Attr("content")
		if !ok {
			continue
		}
		if title != "" {
			return strings.TrimSpace(stringutil.AlphaNumericWithSymbols(title))
		}
	}
	return ""
}

// Body returns the document <body> with whitespace stripped.
func (d *document) Body() string {
	return stringutil.StripWhiteSpace(d.Doc.Find("body").Text())
}

// SocialImage returns a social image from the document.
// Images can be found from Open Graph or Twitter
// properties.
func (d *document) SocialImage() string {
	variations := []string{
		`meta[property="og:image"]`,
		`meta[property=og:image]`,
		`meta[name="twitter:image"]`,
		`meta[name=twitter:image]`,
		`meta[name="twitter:image:alt"]`,
		`meta[name=twitter:image:alt]`,
	}
	for _, variation := range variations {
		image, ok := d.Doc.Find(variation).Attr("content")
		if !ok {
			continue
		}
		uri, err := httputil.GetAbsoluteURL(d.URL.String(), image)
		if err != nil {
			continue
		}
		return uri
	}
	return ""
}

// H1 loops over the <h1>'s of the page and returns the first
// one that isn't empty.
func (d *document) H1() string {
	h1 := ""
	d.Doc.Find("body h1").Each(func(i int, selection *goquery.Selection) {
		if selection.Text() != "" && h1 == "" {
			h1 = sanitize(selection.Text())
		}
	})
	return h1
}

// H2 loops over the <h2>'s of the page and returns the first
// one that isn't empty.
func (d *document) H2() string {
	h2 := ""
	d.Doc.Find("body h2").Each(func(i int, selection *goquery.Selection) {
		if selection.Text() != "" && h2 == "" {
			h2 = sanitize(selection.Text())
		}
	})
	return h2
}

// sanitize sanitises a string by stripping alpha nums and
// trimming white space.
func sanitize(str string) string {
	str = strings.ReplaceAll(str, "—", " ")
	str = strings.ReplaceAll(str, "-", " ")
	return strings.TrimSpace(stringutil.AlphaNumericWithSymbols(str))
}
