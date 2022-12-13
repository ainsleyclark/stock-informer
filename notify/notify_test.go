// Copyright 2023 Ainsley Clark. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package notify

import (
	"context"
	"github.com/ainsleyclark/errors"
	"github.com/ainsleyclark/stock-informer/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	got := New(&config.Config{
		Pages: nil,
		Notify: config.Notify{
			Email: &config.Email{},
			Slack: &config.Slack{},
		},
	})
	assert.NotNil(t, got)
	assert.NotNil(t, got.notify)
}

func TestClient_Send(t *testing.T) {
	tt := map[string]struct {
		send sendFunc
		want any
	}{
		"OK": {
			func(ctx context.Context, subject, message string) error {
				return nil
			},
			nil,
		},
		"Error": {
			func(ctx context.Context, subject, message string) error {
				return errors.New("error")
			},
			"Error notifying element change",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			c := Client{sendFunc: test.send}
			err := c.Send("url", "prev", "now")
			if err != nil {
				assert.Contains(t, errors.Message(err), test.want)
				return
			}
		})
	}
}
