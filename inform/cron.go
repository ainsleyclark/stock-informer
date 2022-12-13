// Copyright 2023 Ainsley Clark. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inform

import (
	"github.com/ainsleyclark/errors"
	"github.com/ainsleyclark/logger"
	"github.com/ainsleyclark/stock-informer/cache"
	"github.com/ainsleyclark/stock-informer/config"
	"github.com/ainsleyclark/stock-informer/crawl"
	"github.com/ainsleyclark/stock-informer/notify"
	"github.com/go-co-op/gocron"
	"strings"
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
		startFunc func()
	}
)

// New instantiates a new cron job.
func New(cfg *config.Config) *Cron {
	schedule := gocron.NewScheduler(time.UTC)
	return &Cron{
		scheduler: schedule,
		cache:     cache.New(),
		config:    cfg,
		scraper:   crawl.New(),
		notifier:  notify.New(cfg),
		startFunc: schedule.StartBlocking,
	}
}

// Boot starts the scheduler to start listening
// and scraping the page.
func (c *Cron) Boot() {
	const op = "Inform.Boot"
	for _, page := range c.config.Pages {
		_, err := c.scheduler.Cron(page.Schedule).Do(c.monitor, page)
		if err != nil {
			logger.WithError(errors.NewInvalid(err, "Error setting up cron job", op)).Error()
		}
	}
	c.startFunc()
}

// Monitor detects a webpage and notifies the user
// of any changes.
func (c *Cron) monitor(page config.Page) {
	// Go and scrape the page and obtain the selector with
	// the relevant selector.
	logger.Debug("Sending request to: " + page.URL)
	element, err := c.scraper.Scrape(page.URL, page.Selector)
	if err != nil {
		logger.WithError(err).Error()
		return
	}

	// Retrieve the item in the cache.
	item, ok := c.cache.Get(page.URL)
	if !ok {
		logger.Debug("No cache item found with URL: " + page.URL)
		// TODO: Recursively call this function as there is no cache.
		c.cache.Set(page.URL, element, cache.RememberForever)
		return
	}

	// Cast to string
	prev := item.(string)

	// If the element stored in the cache is not different
	// to the one we have just crawled, bail.
	formatted := strings.TrimSpace(element)
	if prev == element {
		logger.Debug("No change found for URL: " + page.URL + ", for element: " + formatted)
		return
	}

	// Notify, the element has changed.
	logger.Info("Element changed for URL: " + page.URL + ", for element: " + formatted + ", sending message.")
	err = c.notifier.Send(page.URL, prev, element)
	if err != nil {
		logger.WithError(err).Error()
		return
	}

	// Update the cache.
	c.cache.Set(page.URL, element, cache.RememberForever)
}
