// Created by Yakka (https://theyakka.com)
//
// Copyright (c) 2020 Yakka LLC.
// All rights reserved.
// See the LICENSE file for licensing details and requirements.

package ystore

//type ProcessElementFunc func(key string, val interface{}) interface{}
//
//func (ds *Store) Process(processor ProcessElementFunc) *Store {
//	if ds.Len() == 0 {
//		return NewStore()
//	}
//	processedStore := NewStore()
//	ds.Each(func(k string, v interface{}) {
//		switch v.(type) {
//		case Store:
//			store := v.(Store)
//			sp := &store
//			processedStore.Set(k, sp.Process(processor))
//		case *Store:
//			processedStore.Set(k, v.(*Store).Process(processor))
//		default:
//			processedVal := processor(k, v)
//			if processedVal != nil {
//				processedStore.Set(k, processedVal)
//			}
//		}
//	})
//	return processedStore
//}
//
//type EachFunc func(string, interface{})
//
//func (ds *Store) Each(eachFunc EachFunc) {
//	for k, v := range ds.data {
//		eachFunc(k, v)
//	}
//}
