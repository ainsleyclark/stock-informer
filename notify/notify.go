// Copyright 2023 Ainsley Clark. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package notify

import "github.com/nikoksr/notify"

type (
	Notifier interface {
	}
	Client struct {
		notifier []notify.Notifier
	}
)

func New() *Client {
	return nil
}
