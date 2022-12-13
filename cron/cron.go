// Copyright 2023 Ainsley Clark. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cron

import "github.com/robfig/cron/v3"

type (
	Cron struct {
		client cron.Cron
	}
)
