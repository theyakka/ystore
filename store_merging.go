package ystore

func Merge(store *Store, storesToMerge ...*Store) error {

	// add the entries from the other stores to the primary store
	for _, sm := range storesToMerge {
		for k, e := range sm.entries {
			if err := Set(store, k, e.value); err != nil {
				return err
			}
		}
	}
	return nil
}

func copyEntry(store *Store, entry *Entry, parent *Entry) *Entry {
	copiedEntry := &Entry{
		store:  store,
		key:    entry.key,
		value:  entry.value,
		parent: parent,
	}
	if entry.children != nil {
		copiedEntry.children = map[string]*Entry{}
		for _, child := range entry.children {
			copiedEntry.children[child.key] = copyEntry(store, child, entry)
		}
	}
	return copiedEntry
}
