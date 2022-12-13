// Copyright 2023 Ainsley Clark. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package job

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
	// Cron represents the job that monitors webpages
	// and notifies the user if an element has
	// changed within the DOM.
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
		notifier:  notify.New(cfg),
	}
}

// Boot starts the scheduler to start listening
// and scraping the page.
func (c *Cron) Boot() {
	for _, page := range c.config.Pages {
		_, err := c.scheduler.Cron(page.Schedule).Do(func() {
			c.monitor(page)
		})
		if err != nil {
			log.Fatalln(err)
		}
	}
	c.scheduler.StartBlocking()
}

// Monitor monitors a webpage and notifies the user
// of any changes.
func (c *Cron) monitor(page config.Page) {
	// Go and scrape the page and obtain the selector with
	// the relevant selector.
	element, err := c.scraper.Scrape(page.URL, page.Selector)
	if err != nil {
		log.Println(err)
	}

	// Retrieve the item in the cache
	item, ok := c.cache.Get(page.URL)
	if !ok {
		c.cache.Set(page.URL, element, cache.RememberForever)
		return
	}

	// Cast to string
	compare := item.(string)

	// If the element stored in the cache is not different
	// to the one we have just crawled, bail.
	if compare == element {
		return
	}

	// Notify, the element has changed.
	err = c.notifier.Notify()
	if err != nil {
		log.Println(err)
	}
}
