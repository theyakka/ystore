// Created by Yakka (https://theyakka.com)
//
// Copyright (c) 2020 Yakka LLC.
// All rights reserved.
// See the LICENSE file for licensing details and requirements.

package ystore

import (
	"github.com/spf13/cast"
	"strings"
)

func (ds *Store) Set(key string, value interface{}) {
	splitKey := strings.Split(key, DataKeySeparator)
	if len(splitKey) == 1 {
		ds.data[key] = value
		return
	}
	valMap := ds.GetMapD(splitKey[0], map[string]interface{}{})
	for i := 1; i < len(splitKey)-1; i++ {

	}
	ds.data[splitKey[0]] = valMap
}

func (ds *Store) SetAsString(key string, value interface{}) {
	splitKey := strings.Split(key, DataKeySeparator)
	if len(splitKey) == 1 {
		ds.data[key] = cast.ToString(value)
		return
	}
	var mapValue map[string]interface{}
	for i := len(splitKey) - 1; i > 0; i-- {
		mapValue = map[string]interface{}{
			splitKey[i]: cast.ToString(value),
		}
	}
	ds.data[splitKey[0]] = mapValue

}
