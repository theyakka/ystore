package ystore

import (
	"log"
	"reflect"
)

func FromMap(mapValues map[string]any, options ...Option) *Store {
	store := NewStore(options...)
	AddMapValues(store, store.entries, mapValues)
	return store
}

func AddMapValues(store *Store, entries map[string]*Entry, mapValues map[string]any) {
	for k, v := range mapValues {
		entry := &Entry{
			store: store,
			key:   k,
		}
		switch vt := v.(type) {
		case map[string]any:
			if store.HasFlag(ParseObjects) {
				entry.children = map[string]*Entry{}
				AddMapValues(store, entry.children, v.(map[string]any))
			} else {
				entry.value = reflect.ValueOf(v)
			}
			break
		case []any:
			entry.value = reflect.ValueOf(v)
			break
		default:
			log.Println(vt)
			entry.value = reflect.ValueOf(v)
		}
		entries[k] = entry
	}
}
