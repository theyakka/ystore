// Created by Yakka (https://theyakka.com)
//
// Copyright (c) 2020 Yakka LLC.
// All rights reserved.
// See the LICENSE file for licensing details and requirements.

package ystore

import (
	"github.com/spf13/cast"
)

func (ds *Store) Get(key string) interface{} {
	return ds.GetD(key, nil)
}

func (ds *Store) GetD(key string, defaultValue interface{}) interface{} {
	val := ds.Search(key)
	if val == nil {
		return defaultValue
	}
	return val
}

func (ds *Store) GetString(key string) string {
	return cast.ToString(ds.Get(key))
}

func (ds *Store) GetStringD(key string, defaultValue string) string {
	return cast.ToString(ds.GetD(key, defaultValue))
}

func (ds *Store) GetBool(key string) bool {
	return cast.ToBool(ds.Get(key))
}

func (ds *Store) GetBoolD(key string, defaultValue bool) bool {
	return cast.ToBool(ds.GetD(key, defaultValue))
}

func (ds *Store) GetInt(key string) int {
	return cast.ToInt(ds.Get(key))
}

func (ds *Store) GetIntD(key string, defaultValue int) int {
	return cast.ToInt(ds.GetD(key, defaultValue))
}

func (ds *Store) GetFloat(key string) float64 {
	return cast.ToFloat64(ds.Get(key))
}

func (ds *Store) GetFloatD(key string, defaultValue float64) float64 {
	return cast.ToFloat64(ds.GetD(key, defaultValue))
}

func (ds *Store) GetSlice(key string) []interface{} {
	return cast.ToSlice(ds.Get(key))
}

func (ds *Store) GetSliceD(key string, defaultValue []interface{}) []interface{} {
	return cast.ToSlice(ds.GetD(key, defaultValue))
}

func (ds *Store) GetStringSlice(key string) []string {
	return cast.ToStringSlice(ds.Get(key))
}

func (ds *Store) GetStringSliceD(key string, defaultValue []string) []string {
	return cast.ToStringSlice(ds.GetD(key, defaultValue))
}

func (ds *Store) GetIntSlice(key string) []int {
	return cast.ToIntSlice(ds.Get(key))
}

func (ds *Store) GetIntSliceD(key string, defaultValue []int) []int {
	return cast.ToIntSlice(ds.GetD(key, defaultValue))
}

func (ds *Store) GetMap(key string) map[string]interface{} {
	foundMap := ds.Get(key)
	if foundMap == nil {
		return nil
	}
	return cast.ToStringMap(foundMap)
}

func (ds *Store) GetMapD(key string, defaultValue map[string]interface{}) map[string]interface{} {
	return cast.ToStringMap(ds.GetD(key, defaultValue))
}

func (ds *Store) GetIndexed(key string, index int) interface{} {
	return ds.GetIndexedD(key, index, nil)
}

func (ds *Store) GetIndexedD(key string, index int, defaultValue interface{}) interface{} {
	// TODO - tidy
	val := ds.GetSlice(key)
	if val != nil {
		return val[index]
	}
	return defaultValue
}

func (ds *Store) GetIndexedString(key string, index int) string {
	return ds.GetIndexedStringD(key, index, "")
}

func (ds *Store) GetIndexedStringD(key string, index int, defaultValue string) string {
	if val := ds.GetStringSlice(key); val != nil {
		return val[index]
	}
	return defaultValue
}

func (ds *Store) GetStore(key string) *Store {
	return ds.GetStoreD(key, nil)
}

func (ds *Store) GetStoreD(key string, defaultValue *Store) *Store {
	foundStore := ds.Get(key)
	if foundStore == nil {
		return defaultValue
	}
	if store, ok := foundStore.(*Store); ok {
		return store
	}
	return defaultValue
}
