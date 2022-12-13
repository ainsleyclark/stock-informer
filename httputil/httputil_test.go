// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package httputil

import (
	"github.com/ainsleyclark/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	// badURL is a dirty URL used for testing.
	badURL = "@@@£$%££"
)

func TestIs2xx(t *testing.T) {
	tt := map[string]struct {
		input int
		want  bool
	}{
		"200": {
			200,
			true,
		},
		"201": {
			201,
			true,
		},
		"300": {
			300,
			false,
		},
		"299": {
			200,
			true,
		},
		"199": {
			199,
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := Is2xx(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestIs3xx(t *testing.T) {
	tt := map[string]struct {
		input int
		want  bool
	}{
		"300": {
			300,
			true,
		},
		"301": {
			301,
			true,
		},
		"400": {
			400,
			false,
		},
		"399": {
			399,
			true,
		},
		"299": {
			299,
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := Is3xx(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestGetAbsoluteURL(t *testing.T) {
	tt := map[string]struct {
		fullURL  string
		relative string
		want     any
	}{
		"Full URL Error": {
			badURL,
			"",
			"Error parsing full URI",
		},
		"Relative URL Error": {
			"",
			badURL,
			"Error parsing relative URI",
		},
		"Relative": {
			"https://www.comparethemarket.com/",
			"./img/test.jpg",
			"https://www.comparethemarket.com/img/test.jpg",
		},
		"Forward Slash": {
			"https://www.comparethemarket.com/home-insurance/content/green-eco-friendly-renovations",
			"/home-insurance/content/green-eco-friendly-renovations/",
			"https://www.comparethemarket.com/home-insurance/content/green-eco-friendly-renovations/",
		},
		"Absolute Full URL": {
			"https://www.comparethemarket.com/home-insurance/content/green-eco-friendly-renovations",
			"https://www.comparethemarket.com/home-insurance/content/green-eco-friendly-renovations/",
			"https://www.comparethemarket.com/home-insurance/content/green-eco-friendly-renovations/",
		},
		"Absolute Full URL 2": {
			"https://www.autotrader.co.uk/cars/electric/ev-drivers-with-disabilities",
			"/cars/electric/ev-drivers-with-disabilities/",
			"https://www.autotrader.co.uk/cars/electric/ev-drivers-with-disabilities/",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := GetAbsoluteURL(test.fullURL, test.relative)
			if err != nil {
				assert.Contains(t, errors.Message(err), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}
