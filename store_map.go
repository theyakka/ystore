package ystore

import (
	"reflect"
)

func FromMap(mapValues map[string]any, options ...Option) *Store {
	store := NewStore(options...)
	AddMapValues(store, store.entries, mapValues)
	return store
}

func AddMapValues(store *Store, entries EntriesMap, mapValues map[string]any) {
	for k, v := range mapValues {
		entry := &Entry{
			store: store,
			key:   k,
		}
		switch converted := v.(type) {
		case map[string]any:
			if store.HasFlag(ParseObjects) {
				entry.children = EntriesMap{}
				AddMapValues(store, entry.children, converted)
			} else {
				entry.value = reflect.ValueOf(v)
			}
		case []any:
			entry.value = reflect.ValueOf(v)
		default:
			entry.value = reflect.ValueOf(v)
		}
		entries[k] = entry
	}
}
