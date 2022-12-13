// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package httputil

import (
	"github.com/Danny-Dasilva/CycleTLS/cycletls"
	"github.com/ainsleyclark/errors"
	mocks "github.com/krang-backlink/api/gen/mocks/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func TestNewClient(t *testing.T) {
	got := NewClient()
	assert.NotNil(t, got)
}

func TestHttpClient_Do(t *testing.T) {
	url := "https://google.com"

	tt := map[string]struct {
		mock func(m *mocks.CycleTLS)
		want any
	}{
		"Error": {
			func(m *mocks.CycleTLS) {
				m.On("Do", mock.Anything, mock.Anything, mock.Anything).
					Return(cycletls.Response{}, errors.New("error"))
			},
			"Error performing client request",
		},
		"Redirect": {
			func(m *mocks.CycleTLS) {
				m.On("Do", mock.Anything, mock.Anything, mock.Anything).
					Return(cycletls.Response{Status: http.StatusMovedPermanently, Headers: map[string]string{"Location": "test"}}, nil).
					Once()
				m.On("Do", mock.Anything, mock.Anything, mock.Anything).
					Return(cycletls.Response{Body: "test"}, nil).
					Once()
			},
			"test",
		},
		"Location Error": {
			func(m *mocks.CycleTLS) {
				m.On("Do", mock.Anything, mock.Anything, mock.Anything).
					Return(cycletls.Response{Status: http.StatusMovedPermanently, Headers: map[string]string{"Location": badURL}}, nil)
			},
			"Error parsing relative URI",
		},
		"Same Location": {
			func(m *mocks.CycleTLS) {
				m.On("Do", mock.Anything, mock.Anything, mock.Anything).
					Return(cycletls.Response{Status: http.StatusMovedPermanently, Headers: map[string]string{"Location": url}, Body: "test"}, nil)
			},
			"test",
		},
		"OK": {
			func(m *mocks.CycleTLS) {
				m.On("Do", mock.Anything, mock.Anything, mock.Anything).
					Return(cycletls.Response{Body: "test"}, nil)
			},
			"test",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			cycle := &mocks.CycleTLS{}
			if test.mock != nil {
				test.mock(cycle)
			}
			client := httpClient{cycle: cycle}
			_, err := client.Do(url, http.MethodGet)
			if err != nil {
				assert.Contains(t, errors.Message(err), test.want)
				return
			}
		})
	}
}
