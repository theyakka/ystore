package ystore

import "strings"

func (ds Store) Set(key string, value interface{}) {
	splitKey := strings.Split(key, DataKeySeparator)
	if len(splitKey) == 1 {
		ds.data[key] = value
		return
	}
	mapValue := value
	for i := len(splitKey) - 1; i > 0; i-- {
		mapValue = map[string]interface{}{
			splitKey[i]: mapValue,
		}
	}
	ds.data[splitKey[0]] = mapValue
}
