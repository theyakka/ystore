package ystore

import (
	"reflect"
	"strings"
)

func Set(store *Store, keyPath string, value any) error {
	return setValue(store, store.entries, keyPath, value)
}

func (s *Store) Set(keyPath string, value any) error {
	return Set(s, keyPath, value)
}

func setValue(store *Store, entries EntriesMap, keyPath string, value any) error {
	// check if it's cached and, if so, set the value of the cached item
	entry := store.Get(keyPath)
	if entry != nil {
		entry.value = reflect.ValueOf(value)
		return nil
	}
	// split the key path and see if it has only one segment. if so,
	// we can set the item in the entries map quickly and be done
	segments := strings.Split(keyPath, ".")
	if len(segments) == 1 {
		entries[segments[0]] = &Entry{
			store: store,
			key:   segments[0],
			value: reflect.ValueOf(value),
		}
		return nil
	}
	// find the entry (as it is not cached) or create it as we search for it
	_ = findOrInsertEntry(store, entries, segments, value, nil)
	// check to see if a driver has been assigned to the store, that the driver has
	// parameters set and if the driver supports auto persistence. If so, we should
	// persist the data here.
	if d := store.driver; d != nil && d.Parameters() != nil && d.Parameters().AutoPersist {
		return d.Persist()
	}
	return nil
}

func findOrInsertEntry(store *Store, entries EntriesMap, pathSegments []string, value any, parent *Entry) *Entry {
	entry := FindEntry(entries, pathSegments)
	if entry != nil {
		entry.value = reflect.ValueOf(value)
	}

	segment := pathSegments[0]
	entry = entries[segment]
	if entry == nil {
		entry = &Entry{
			store:  store,
			key:    segment,
			parent: parent,
		}
		entries[segment] = entry
	}

	if len(pathSegments) == 1 {
		entry.value = reflect.ValueOf(value)
		return entry
	}

	if entry.children == nil {
		entry.children = EntriesMap{}
	}
	return findOrInsertEntry(store, entry.children, pathSegments[1:], value, entry)
}
