// Copyright 2023 Ainsley Clark. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package notify

import (
	"context"
	"github.com/ainsleyclark/stock-informer/config"
	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/mail"
)

type (
	Notifier interface {
		Notify() error
	}
	Client struct {
		cfg *config.Config
	}
)

func New(cfg *config.Config) *Client {
	c := &Client{
		cfg: cfg,
	}
	if cfg.Notify.Email != nil {
		notify.UseServices(c.email())
	}
	return c
}

func (c *Client) Notify() error {
	// Send a test message.
	_ = notify.Send(
		context.Background(),
		"Subject/Title",
		"The actual message - Hello, you awesome gophers! :)",
	)
	return nil
}

func (c *Client) email() notify.Notifier {
	cfg := c.cfg.Notify.Email
	m := mail.New(cfg.Sender, cfg.Address)
	m.AuthenticateSMTP("", cfg.User, cfg.Password, cfg.Address)
	m.AddReceivers(cfg.Receivers...)
	return m
}
