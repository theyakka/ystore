// Created by Yakka (https://theyakka.com)
//
// Copyright (c) 2022 Yakka LLC.
// All rights reserved.
// See the LICENSE file for licensing details and requirements.

package ystore

import (
	"fmt"
	"sync"
)

const (
	defaultCacheState = true
	defaultCacheSize  = 50
)

type EntriesMap map[string]*Entry

type Store struct {
	driver      Driver
	flags       Flag
	entries     EntriesMap
	enableCache bool
	cacheLimit  int
	cache       EntriesMap
	mutex       sync.RWMutex
}

func NewStore(options ...Option) *Store {
	store := &Store{
		flags:       ParseObjects,
		entries:     EntriesMap{},
		enableCache: defaultCacheState,
		cacheLimit:  defaultCacheSize,
	}
	for _, opt := range options {
		opt(store)
	}
	if store.enableCache {
		// only initialize the cache storage if we have enabled the cache
		store.cache = EntriesMap{}
	}
	return store
}

func (s *Store) SetDriver(driver Driver) {
	s.driver = driver
}

func (s *Store) Load(uris ...string) error {
	return s.driver.Load(s, uris...)
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
	s.entries = EntriesMap{}
}

func (s *Store) Debug() {
	printEntries(s.Entries())
}

func printEntries(entries EntriesMap) {
	for _, e := range entries {
		fmt.Println(e.Key(), "=", e.Value(), ":", e.Kind())
		if e.children != nil {
			printEntries(e.children)
		}
	}
}

type Mutable interface {
	Set(keyPath string, value any) error
}

type Readable interface {
	Get(keyPath string) Readable
}
