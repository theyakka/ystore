package ystore

func (e *Entry) Set(keyPath string, value any) error {
	// if the entry has no children then we need to make the empty holder
	// to add whatever is being requested.
	if e.children == nil {
		e.children = EntriesMap{}
	}
	return setValue(e.store, e.children, keyPath, value)
}
