// Copyright 2023 Ainsley Clark. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"github.com/ainsleyclark/logger"
	"github.com/ainsleyclark/stock-informer/config"
	"github.com/ainsleyclark/stock-informer/job"
	"log"
)

func main() {
	// Create the logger.
	opts := logger.NewOptions().
		Service("Stock Informer").
		Prefix("INFORMER").
		DefaultStatus("LOG")
	err := logger.New(context.Background(), opts)
	if err != nil {
		log.Fatalln(err)
	}

	// Load the configuration file.
	logger.Info("Loading Configuration")
	cfg, err := config.Load("/Users/ainsley/Desktop/Web/apis/stock-informer/config.yml")
	if err != nil {
		logger.Fatal(err)
	}

	// Boot the cron job.
	logger.Info("Booting Informer")
	job.New(cfg).Boot()
}
