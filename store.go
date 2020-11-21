// Created by Yakka (https://theyakka.com)
//
// Copyright (c) 2020 Yakka LLC.
// All rights reserved.
// See the LICENSE file for licensing details and requirements.

package ystore

// Store is a giant data map that is constructed from one or more data files
// that are stored within a directory and series of sub-directories
type Store struct {
	// data is the primary storage for all the parsed data/config files
	data map[string]interface{}
	// options that we want to configure for the store
	options StoreOptions
}

//
func NewStore(options StoreOptions) *Store {
	return &Store{
		data:    map[string]interface{}{},
		options: options,
	}
}

// NewStoreFromFile will create a new store and then add the contents of the file to the store.
// If an error occurs while loading/parsing the file, then the error will be ignored and an
// empty store will be returned.
func NewStoreFromFile(filename string, options StoreOptions) *Store {
	store := NewStore(options)
	_ = store.AddFile(filename)
	return store
}

func (ds *Store) AllValues() map[string]interface{} {
	return ds.data
}

func (ds *Store) Len() int {
	return len(ds.data)
}
