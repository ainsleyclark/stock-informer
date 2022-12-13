// Copyright 2023 Ainsley Clark. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/ainsleyclark/errors"
	"github.com/ainsleyclark/logger"
	"github.com/ainsleyclark/stock-informer/config"
	"github.com/ainsleyclark/stock-informer/inform"
	"github.com/enescakir/emoji"
	"log"
)

func main() {
	// Info.
	fmt.Printf("\n%v Welcome to Stock Informer\n\n", emoji.WavingHand)

	// Create the logger.
	opts := logger.NewOptions().
		Service("Stock Informer").
		Prefix("INFORMER").
		DefaultStatus("LOG")
	err := logger.New(context.Background(), opts)
	if err != nil {
		log.Fatalln(err)
	}

	// Obtain path flag.
	path := flag.String("path", "", "Configuration file path")
	flag.Parse()
	if path == nil || *path == "" {
		logger.WithError(errors.NewInvalid(errors.New("no path detected"), "Error, no path found, use -path=./config.yml for the configuration file.", "Configuration.Validate")).Fatal()
	}

	// Load the configuration file.
	logger.Info("Loading Configuration")
	cfg, err := config.Load("/Users/ainsley/Desktop/Web/apis/stock-informer/config.yml")
	if err != nil {
		logger.Fatal(err)
	}

	// Validate config.
	// TODO - Move to separate config validation function.
	if cfg.Pages == nil {
		logger.WithError(errors.NewInvalid(errors.New("invalid configuration"), "Error, no pages found to scrape", "Configuration.Validate")).Fatal()
	}

	// Boot the cron job.
	logger.Info("Booting Informer")
	inform.New(cfg).Boot()
}
