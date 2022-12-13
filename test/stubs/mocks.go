// Copyright 2023 Ainsley Clark. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package stubs

import "github.com/Danny-Dasilva/CycleTLS/cycletls"

type CycleTLS interface {
	Do(string, cycletls.Options, string) (cycletls.Response, error)
}
