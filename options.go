// Created by Yakka (https://theyakka.com)
//
// Copyright (c) 2022 Yakka LLC.
// All rights reserved.
// See the LICENSE file for licensing details and requirements.

package ystore

type Option func(store *Store)

func WithCache(enabled bool, size int) Option {
	return func(store *Store) {
		store.enableCache = enabled
		store.cacheLimit = size
	}
}
