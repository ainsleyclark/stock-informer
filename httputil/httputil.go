// Copyright 2023 Ainsley Clark. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package httputil

import (
	"github.com/ainsleyclark/errors"
	"net/url"
	"strings"
)

// Is2xx determines if a response status code is flagged as OK.
func Is2xx(status int) bool {
	if status < 200 || status >= 300 {
		return false
	}
	return true
}

// Is3xx determines if a response status code is a redirect.
func Is3xx(status int) bool {
	if status < 300 || status >= 400 {
		return false
	}
	return true
}

// GetAbsoluteURL retrieves the absolute URL of a full and
// relative URL.
// Returns errors.INVALID if the urls could not be parsed.
func GetAbsoluteURL(fullURL string, relative string) (string, error) {
	const op = "HTTPUtil.GetAbsoluteURL"
	full, err := url.Parse(fullURL)
	if err != nil {
		return "", errors.NewInvalid(err, "Error parsing full URI", op)
	}
	rel, err := url.Parse(relative)
	if err != nil {
		return "", errors.NewInvalid(err, "Error parsing relative URI", op)
	}
	if !strings.Contains(relative, "http") && !strings.HasPrefix(relative, "./") {
		return full.Scheme + "://" + full.Host + "/" + strings.TrimPrefix(relative, "/"), nil
	}
	if rel.IsAbs() {
		return relative, nil
	}
	return strings.TrimSuffix(fullURL, "/") + "/" + strings.TrimPrefix(strings.TrimPrefix(relative, "./"), "/"), nil
}
