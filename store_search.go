package ystore

import (
	"reflect"
	"strings"
)

const SearchPathDelimiter = "."

type SearchOption func(options *SearchOptions)

type SearchOptions struct {
	searchStructs bool
}

func SearchStore(store *Store, keyPath string, options SearchOptions) interface{} {
	if keyPath == "" {
		return nil
	}
	splitPath := strings.Split(keyPath, SearchPathDelimiter)
	if len(splitPath) == 1 {
		return store.Get(splitPath[0])
	}
	return searchPaths(store.data, splitPath, options)
}

func searchPaths(source interface{}, paths []string, options SearchOptions) interface{} {
	if len(paths) == 0 || source == nil {
		return source // done
	}

	switch source.(type) {

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

}
