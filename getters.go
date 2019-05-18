package ystore

import (
	"strings"

	"github.com/spf13/cast"
)

func (ds Store) Get(key string) interface{} {
	splitKey := strings.Split(key, DataKeySeparator)
	val := SearchMap(ds.data, splitKey)
	if val == nil {
		return nil
	}
	return val
}

func (ds Store) GetD(key string, defaultValue interface{}) interface{} {
	splitKey := strings.Split(key, DataKeySeparator)
	val := SearchMap(ds.data, splitKey)
	if val == nil {
		return defaultValue
	}
	return val
}

func (ds Store) GetString(key string) string {
	return cast.ToString(ds.Get(key))
}

func (ds Store) GetStringD(key string, defaultValue string) string {
	return cast.ToString(ds.GetD(key, defaultValue))
}

func (ds Store) GetBool(key string) bool {
	return cast.ToBool(ds.Get(key))
}

func (ds Store) GetBoolD(key string, defaultValue bool) bool {
	return cast.ToBool(ds.GetD(key, defaultValue))
}

func (ds Store) GetInt(key string) int {
	return cast.ToInt(ds.Get(key))
}

func (ds Store) GetIntD(key string, defaultValue int) int {
	return cast.ToInt(ds.GetD(key, defaultValue))
}

func (ds Store) GetFloat(key string) float64 {
	return cast.ToFloat64(ds.Get(key))
}

func (ds Store) GetFloatD(key string, defaultValue float64) float64 {
	return cast.ToFloat64(ds.GetD(key, defaultValue))
}

func (ds Store) GetStringSlice(key string) []string {
	return cast.ToStringSlice(ds.Get(key))
}

func (ds Store) GetStringSliceD(key string, defaultValue []string) []string {
	return cast.ToStringSlice(ds.GetD(key, defaultValue))
}

func (ds Store) GetIntSlice(key string) []int {
	return cast.ToIntSlice(ds.Get(key))
}

func (ds Store) GetIntSliceD(key string, defaultValue []int) []int {
	return cast.ToIntSlice(ds.GetD(key, defaultValue))
}

func (ds Store) GetMap(key string) map[string]interface{} {
	foundMap := ds.Get(key)
	if foundMap == nil {
		return nil
	}
	return cast.ToStringMap(foundMap)
}

func (ds Store) GetMapD(key string, defaultValue map[string]interface{}) map[string]interface{} {
	return cast.ToStringMap(ds.GetD(key, map[string]interface{}{}))
}
