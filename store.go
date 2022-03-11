// Created by Yakka (https://theyakka.com)
//
// Copyright (c) 2022 Yakka LLC.
// All rights reserved.
// See the LICENSE file for licensing details and requirements.

package ystore

import (
	"sync"
)

const (
	defaultCacheState = true
	defaultCacheSize  = 50
)

type Store struct {
	driver      Driver
	flags       Flag
	entries     map[string]*Entry
	enableCache bool
	cacheLimit  int
	cache       map[string]*Entry
	mutex       sync.RWMutex
}

func NewStore(options ...Option) *Store {
	store := &Store{
		flags:       ParseObjects,
		entries:     map[string]*Entry{},
		enableCache: defaultCacheState,
		cacheLimit:  defaultCacheSize,
	}
	for _, opt := range options {
		opt(store)
	}
	if store.enableCache {
		// only initialize the cache storage if we have enabled the cache
		store.cache = map[string]*Entry{}
	}
	return store
}

func (s *Store) SetDriver(driver Driver) {
	s.driver = driver
}

func (s *Store) Load(uri string) error {
	return s.driver.Load(s, uri)
}

func (s *Store) Persist() error {
	if s.driver == nil {
		return nil
	}
	return s.driver.Persist()
}

func (s *Store) HasFlag(flag Flag) bool {
	return s.flags&flag != 0
}

func (s *Store) Clear() {
	s.entries = map[string]*Entry{}
}
