// Copyright 2023 Ainsley Clark. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/ainsleyclark/errors"
	"gopkg.in/yaml.v3"
	"os"
)

type (
	// Config represents the configuration for the application.
	Config struct {
		URLs   []URL  `yaml:"urls"`
		Notify Notify `yaml:"notify"`
	}
	// URL represents a singular URL to monitor.
	URL struct {
		URL      string `yaml:"url"`
		Selector string `yaml:"selector"`
		Schedule string `yaml:"schedule"`
	}
	// Notify represents the notification settings for when
	// an element has changed in the DOM.
	Notify struct {
		Email Email `yaml:"email"`
		Slack Slack `yaml:"slack"`
	}
	// Email represents SMTP email credentials.
	Email struct {
		Address  string `yaml:"address"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	}
	// Slack represents the Slack credentials configuration.
	Slack struct {
		Token     string `yaml:"token"`
		ChannelID string `yaml:"channel_id"`
	}
)

// Load generates and loads the configuration by
// a specified path.
// Returns an error if the lookup failed or the yaml
// could not be unmarshalled.
func Load(path string) (Config, error) {
	const op = "Configuration.Load"
	bytes, err := os.ReadFile(path)
	if err != nil {
		return Config{}, errors.NewInvalid(err, "Error finding configuration path with the path: "+path, op)
	}
	config := Config{}
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return Config{}, errors.NewInvalid(err, "Error loading configuration", op)
	}
	return config, nil
}
