package ystore

func Merge(storesToMerge []*Store, options ...Option) (*Store, error) {
	// create the store that will hold all of our merged entries
	// if you don't pass any stores to merge then we don't have any
	// work, so we can just be done. in this case we always return an
	// empty store
	if storesToMerge == nil || len(storesToMerge) == 0 {
		return NewStore(options...), nil
	}
	mergedStore := NewStore(options...)
	// add the entries from the other stores to the primary store
	for _, sm := range storesToMerge {
		if err := mergeEntries(mergedStore, sm.entries, nil); err != nil {
			return NewStore(options...), err
		}
	}
	return mergedStore, nil
}

func (s *Store) Merge(storesToMerge []*Store, options ...Option) (*Store, error) {
	stores := []*Store{s}
	stores = append(stores, storesToMerge...)
	return Merge(stores, options...)
}

func mergeEntries(store *Store, entries EntriesMap, parent *Entry) error {
	if entries == nil || len(entries) == 0 {
		return nil
	}
	for k, e := range entries {
		var dest Mutable = store
		entry := store.Get(e.KeyPath())
		if parent != nil {
			entry = parent.Get(k)
			dest = parent
		}
		if entry == nil {
			entry = &Entry{
				store:    store,
				key:      k,
				parent:   parent,
				children: nil,
			}
		}
		if e.HasValue() {
			entry.value = e.value
		}
		_ = dest.Set(k, e)
		if e.HasChildren() {
			_ = mergeEntries(store, e.children, entry)
		}
	}
	return nil
}
