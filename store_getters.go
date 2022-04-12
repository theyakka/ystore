// Created by Yakka (https://theyakka.com)
//
// Copyright (c) 2022 Yakka LLC.
// All rights reserved.
// See the LICENSE file for licensing details and requirements.

package ystore

import (
	"strings"
)

func (s *Store) Entries() EntriesMap {
	return s.entries
}

func (s *Store) Has(keyPath string) bool {
	s.mutex.RLock()
	entry := FindEntry(s, s.entries, strings.Split(keyPath, "."))
	s.mutex.RUnlock()
	return entry != nil
}

func Get(store *Store, keyPath string) *Entry {
	store.mutex.RLock()
	entry := store.cache[keyPath]
	if entry != nil {
		store.mutex.RUnlock()
		return entry
	}
	entry = FindEntry(store, store.entries, strings.Split(keyPath, "."))
	if store.enableCache && entry != nil {
		store.cache[keyPath] = entry
	}
	store.mutex.RUnlock()
	return entry
}

func (s *Store) Get(keyPath string) *Entry {
	return Get(s, keyPath)
}

func FindEntry(store *Store, entries EntriesMap, pathSegments []string) *Entry {
	segment := pathSegments[0]
	segmentCount := len(pathSegments)
	entry := entries[segment]
	if entry == nil || segmentCount == 1 {
		return entry
	}
	return FindEntry(store, entry.children, pathSegments[1:])
}
