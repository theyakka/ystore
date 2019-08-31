package ystore

type ProcessElementFunc func(key string, val interface{}) interface{}

func (ds Store) Process(processor ProcessElementFunc) *Store {
	if ds.Len() == 0 {
		return NewStore()
	}
	processedStore := NewStore()
	ds.Each(func(k string, v interface{}) {
		switch v.(type) {
		case Store:
			processedStore.Set(k, v.(Store).Process(processor))
		case *Store:
			processedStore.Set(k, v.(*Store).Process(processor))
		default:
			processedVal := processor(k, v)
			if processedVal != nil {
				processedStore.Set(k, processedVal)
			}
		}
	})
	return processedStore
}

type EachFunc func(string, interface{})

func (ds Store) Each(eachFunc EachFunc) {
	for k, v := range ds.data {
		eachFunc(k, v)
	}
}
