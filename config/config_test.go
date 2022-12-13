// Copyright 2023 Ainsley Clark. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/ainsleyclark/errors"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)

	tt := map[string]struct {
		input string
		want  any
	}{
		"OK": {
			filepath.Join(wd, "testdata", "config.yml"),
			Config{
				URLs: []URL{
					{URL: "https://test.com", Selector: ".div", Schedule: "0 30 * * * *"},
				},
				Notify: Notify{
					Email: Email{Address: "smtp.gmail.com", User: "hello@hello.com", Password: "password"},
					Slack: Slack{Token: "token", ChannelID: "id"},
				},
			},
		},
		"Bath Path": {
			"wrong",
			"Error finding configuration path with the path",
		},
		"Unmarshal Error": {
			filepath.Join(wd, "testdata", "config-invalid.yml"),
			"Error loading configuration",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := Load(test.input)
			if err != nil {
				assert.Contains(t, errors.Message(err), test.want)
				return
			}
			assert.Equal(t, test.want, *got)
		})
	}
}
