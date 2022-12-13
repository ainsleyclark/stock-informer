// Copyright 2023 Ainsley Clark. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/ainsleyclark/errors"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type (
	// Config represents the configuration for the application.
	Config []MonitorURL
	// MonitorURL represents a singular URL to monitor.
	MonitorURL struct {
		URL         string `yaml:"url"`
		Path        string `yaml:"path"`
		MonitorTime string `yaml:"monitor-time"`
	}
)

// Load generates and loads the configuration by
// a specified path.
// Returns an error if the lookup failed or the yaml
// could not be unmarshalled.
func Load(path string) (Config, error) {
	filePath := filepath.Join(path, "config.yml")
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, errors.NewInvalid(err, "Error finding configuration path with the path: ", filePath)
	}
	config := Config{}
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
