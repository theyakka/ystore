package ystore

func Copy(store *Store, options ...Option) *Store {
	copiedStore := NewStore(options...)
	for k, e := range store.entries {
		copiedStore.entries[k] = copyEntry(store, e, nil)
	}
	return copiedStore
}

func (s *Store) Copy() *Store {
	return Copy(s)
}

func copyEntry(store *Store, entry *Entry, parent *Entry) *Entry {
	copiedEntry := &Entry{
		store:  store,
		key:    entry.key,
		value:  entry.value,
		parent: parent,
	}
	if entry.children != nil {
		copiedEntry.children = EntriesMap{}
		for _, child := range entry.children {
			copiedEntry.children[child.key] = copyEntry(store, child, entry)
		}
	}
	return copiedEntry
}
