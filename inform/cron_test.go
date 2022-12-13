// Copyright 2023 Ainsley Clark. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inform

import (
	"bytes"
	"github.com/ainsleyclark/errors"
	"github.com/ainsleyclark/logger"
	"github.com/ainsleyclark/stock-informer/config"
	cache "github.com/ainsleyclark/stock-informer/mocks/cache"
	crawl "github.com/ainsleyclark/stock-informer/mocks/crawl"
	notify "github.com/ainsleyclark/stock-informer/mocks/notify"
	"github.com/go-co-op/gocron"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	got := New(&config.Config{})
	assert.NotNil(t, got.scheduler)
	assert.NotNil(t, got.cache)
	assert.NotNil(t, got.config)
	assert.NotNil(t, got.scraper)
	assert.NotNil(t, got.notifier)
}

func TestCron_Boot(t *testing.T) {
	tt := map[string]struct {
		input []config.Page
		want  any
	}{
		"OK": {
			[]config.Page{{Schedule: "* * * * *"}},
			"",
		},
		"Cron Error": {
			[]config.Page{{Schedule: "wrong"}},
			"Error setting up cron job",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			c := &Cron{
				scheduler: gocron.NewScheduler(time.UTC),
				config:    &config.Config{Pages: test.input},
				startFunc: func() {},
			}
			buf := &bytes.Buffer{}
			logger.SetLevel(logrus.TraceLevel)
			logger.SetOutput(buf)
			c.Boot()
			assert.Contains(t, buf.String(), test.want)
		})
	}
}

func TestCron_Monitor(t *testing.T) {
	page := config.Page{
		URL:      "https://github.com/ainsleyclark",
		Selector: ".test",
	}

	tt := map[string]struct {
		mock func(store *cache.Store, scraper *crawl.Scraper, notifier *notify.Notifier)
		want any
	}{
		"Scrape Error": {
			func(store *cache.Store, scraper *crawl.Scraper, notifier *notify.Notifier) {
				scraper.On("Scrape", page.URL, page.Selector).
					Return("", &errors.Error{Message: "scrape error"})
			},
			"scrape error",
		},
		"Cache Not Set": {
			func(store *cache.Store, scraper *crawl.Scraper, notifier *notify.Notifier) {
				scraper.On("Scrape", page.URL, page.Selector).
					Return("element", nil)
				store.On("Get", page.URL).
					Return("", false)
				store.On("Set", page.URL, "element", mock.Anything)
			},
			"No cache item found with URL",
		},
		"No Change": {
			func(store *cache.Store, scraper *crawl.Scraper, notifier *notify.Notifier) {
				scraper.On("Scrape", page.URL, page.Selector).
					Return("element", nil)
				store.On("Get", page.URL).
					Return("element", true)
			},
			"No change found",
		},
		"Notify Error": {
			func(store *cache.Store, scraper *crawl.Scraper, notifier *notify.Notifier) {
				scraper.On("Scrape", page.URL, page.Selector).
					Return("element", nil)
				store.On("Get", page.URL).
					Return("prev", true)
				notifier.On("Send", page.URL, "prev", "element").
					Return(&errors.Error{Message: "notify error"})
			},
			"notify error",
		},
		"Changed OK": {
			func(store *cache.Store, scraper *crawl.Scraper, notifier *notify.Notifier) {
				scraper.On("Scrape", page.URL, page.Selector).
					Return("element", nil)
				store.On("Get", page.URL).
					Return("prev", true)
				notifier.On("Send", page.URL, "prev", "element").
					Return(nil)
				store.On("Set", page.URL, "element", mock.Anything)
			},
			"Element changed for URL",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			var (
				store    = &cache.Store{}
				scraper  = &crawl.Scraper{}
				notifier = &notify.Notifier{}
			)
			if test.mock != nil {
				test.mock(store, scraper, notifier)
			}
			c := Cron{
				cache:    store,
				scraper:  scraper,
				notifier: notifier,
			}
			buf := &bytes.Buffer{}
			logger.SetLevel(logrus.TraceLevel)
			logger.SetOutput(buf)
			c.monitor(page)
			assert.Contains(t, buf.String(), test.want)
		})
	}
}
