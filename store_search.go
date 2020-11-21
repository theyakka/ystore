package ystore

import (
	"reflect"
	"strings"
)

type SearchOption func(options *SearchOptions)

type SearchOptions struct {
	searchStructs bool
}

func NewSearchOptions() SearchOptions {
	return SearchOptions{
		searchStructs: true,
	}
}

func (ds *Store) Search(keyPath string, options ...SearchOption) interface{} {
	return SearchStore(ds, keyPath, options...)
}

func SearchStore(store *Store, keyPath string, options ...SearchOption) interface{} {
	if keyPath == "" {
		return nil
	}
	splitPath := strings.Split(keyPath, DataKeySeparator)
	if len(splitPath) == 1 {
		return store.data[keyPath]
	}
	return searchPaths(store.data, splitPath, NewSearchOptions())
}

func searchPaths(source interface{}, paths []string, options SearchOptions) interface{} {
	if len(paths) == 0 || source == nil {
		return source // done
	}
	switch source.(type) {
	case Store:
		store := source.(Store)
		return searchPaths(store.data, paths[1:], options)
	case *Store:
		store := source.(*Store)
		return searchPaths(store.data, paths[1:], options)
	}
	v := reflect.ValueOf(source)
	if v.Kind() == reflect.Map {
		mapVal := v.MapIndex(reflect.ValueOf(paths[0]))
		if mapVal.IsZero() {
			return nil
		}
		return searchPaths(mapVal.Interface(), paths[1:], options)
	} else if options.searchStructs && v.Kind() == reflect.Struct {
		field := v.FieldByName(paths[0])
		if field.IsZero() {
			return nil
		}
		return searchPaths(field.Interface(), paths[1:], options)
	}
	return nil
}
