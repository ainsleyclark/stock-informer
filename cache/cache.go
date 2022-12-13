// Copyright 2023 Ainsley Clark. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

// Store defines methods for interacting with the
// redis instance.
type Store interface {
	// Get an item from the cache. Returns the item or nil, and a bool
	// indicating whether the key was found.
	Get(k string) (interface{}, bool)
	// Set adds an item to the cache, replacing any existing item. If the duration is 0
	// (DefaultExpiration), the cache's default expiration time is used. If it is -1
	// (NoExpiration), the item never expires.
	Set(k string, x interface{}, d time.Duration)
	// Delete an item from the cache. Does nothing if the key is not in the cache.
	Delete(k string)
	// Flush deletes all items from the cache.
	Flush()
}

const (
	// RememberForever is for use with functions that take an expiration time.
	RememberForever = cache.NoExpiration
)

// New creates a new cache provider.
func New() Store {
	return cache.New(RememberForever, 10*time.Minute)
}
