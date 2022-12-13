// Copyright 2023 Ainsley Clark. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package notify

import (
	"context"
	"fmt"
	"github.com/ainsleyclark/errors"
	"github.com/ainsleyclark/stock-informer/config"
	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/mail"
	"github.com/nikoksr/notify/service/slack"
)

type (
	// Notifier defines the behavior for notification services.
	// The Send function sends a URL with previous and
	// element that is now on the page.
	Notifier interface {
		// Send sends a notification to the services with
		// a message indicating what URL and element
		// has changed.
		// Returns errors.INTERNAL if the message could not be sent.
		Send(url, prev, now string) error
	}
	// Client represents the data for sending
	// notifications to the user.
	Client struct {
		cfg *config.Config
	}
)

// New instantiates a new Notifier client.
func New(cfg *config.Config) *Client {
	c := &Client{
		cfg: cfg,
	}
	if cfg.Notify.Email != nil {
		notify.UseServices(c.email())
	}
	if cfg.Notify.Slack != nil {
		notify.UseServices(c.slack())
	}
	return c
}

// Send sends a notification to the services with
// a message indicating what URL and element
// has changed.
// Returns errors.INTERNAL if the message could not be sent.
func (c *Client) Send(url, prev, now string) error {
	const op = "Notify.Send"
	err := notify.Send(
		context.Background(),
		"Stock Informer - Element Changed",
		fmt.Sprintf("Webpage changed for URL: %s, Previous contents: %s has changed to: %s", url, prev, now),
	)
	if err != nil {
		return errors.NewInternal(err, "Error notifying element change", op)
	}
	return nil
}

// email registers the Mail notifier.
func (c *Client) email() notify.Notifier {
	cfg := c.cfg.Notify.Email
	m := mail.New(cfg.User, fmt.Sprintf("%s:%s", cfg.Address, cfg.Port))
	m.AuthenticateSMTP("", cfg.User, cfg.Password, cfg.Address)
	m.BodyFormat(mail.PlainText)
	m.AddReceivers(cfg.Receivers...)
	return m
}

// slack registers the Slack notifier.
func (c *Client) slack() notify.Notifier {
	cfg := c.cfg.Notify.Slack
	s := slack.New(cfg.Token)
	s.AddReceivers(cfg.ChannelID)
	return s
}
