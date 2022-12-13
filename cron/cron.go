// Copyright 2023 Ainsley Clark. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cron

import (
	"github.com/ainsleyclark/stock-informer/cache"
	"github.com/ainsleyclark/stock-informer/config"
	"github.com/ainsleyclark/stock-informer/crawl"
	"github.com/ainsleyclark/stock-informer/notify"
	"github.com/go-co-op/gocron"
	"log"
	"time"
)

type (
	Cron struct {
		scheduler *gocron.Scheduler
		cache     cache.Store
		config    *config.Config
		scraper   crawl.Scraper
		notifier  notify.Notifier
	}
)

// New instantiates a new cron job.
func New(cfg *config.Config) *Cron {
	return &Cron{
		scheduler: gocron.NewScheduler(time.UTC),
		cache:     cache.New(),
		config:    cfg,
		scraper:   crawl.New(),
		notifier:  notify.New(),
	}
}

// Boot starts the scheduler to start listening
// and scraping the page.
func (c *Cron) Boot() {
	for _, url := range c.config.URLs {
		_, err := c.scheduler.Cron(url.Schedule).Do(func() {
			scrape, err := c.scraper.Scrape(url.URL, url.Selector)
			if err != nil {
				log.Println(err)
			}

		})
		if err != nil {
			log.Fatalln(err)
		}
	}
}
