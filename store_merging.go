// Created by Yakka (https://theyakka.com)
//
// Copyright (c) 2022 Yakka LLC.
// All rights reserved.
// See the LICENSE file for licensing details and requirements.

package ystore

func (s *Store) Merge(storesToMerge []*Store, options ...Option) (*Store, error) {
	stores := []*Store{s}
	stores = append(stores, storesToMerge...)
	return Merge(stores, options...)
}

func Merge(storesToMerge []*Store, options ...Option) (*Store, error) {
	// create the store that will hold all of our merged entries
	// if you don't pass any stores to merge then we don't have any
	// work, so we can just be done. in this case we always return an
	// empty store
	if len(storesToMerge) == 0 {
		return NewStore(options...), nil
	}
	mergedStore := NewStore(options...)
	// add the entries from the other stores to the primary store
	for _, sm := range storesToMerge {
		if err := mergeEntries(mergedStore, sm.entries); err != nil {
			return NewStore(options...), err
		}
	}
	return mergedStore, nil
}

func mergeEntries(toStore *Store, entries EntriesMap) error {
	if len(entries) == 0 {
		return nil
	}
	for _, e := range entries {
		if e.HasValue() {
			if err := toStore.Set(e.KeyPath(), e.value); err != nil {
				return err
			}
		}
		if e.HasChildren() {
			if err := mergeEntries(toStore, e.children); err != nil {
				return err
			}
		}
	}
	return nil
}
