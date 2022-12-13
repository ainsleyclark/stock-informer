// Copyright 2023 Ainsley Clark. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cron

import (
	"github.com/ainsleyclark/stock-informer/config"
	"github.com/ainsleyclark/stock-informer/crawl"
	"github.com/ainsleyclark/stock-informer/notify"
	"github.com/robfig/cron/v3"
)

type (
	Cron struct {
		client   *cron.Cron
		config   *config.Config
		scraper  crawl.Scraper
		notifier notify.Notifier
	}
)

// New instantiates a new cron job.
func New(cfg *config.Config) *Cron {
	return &Cron{
		client:   cron.New(),
		config:   cfg,
		scraper:  crawl.New(),
		notifier: notify.New(),
	}
}
