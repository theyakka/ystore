// Created by Yakka (https://theyakka.com)
//
// Copyright (c) 2022 Yakka LLC.
// All rights reserved.
// See the LICENSE file for licensing details and requirements.

package ystore_test

import (
	"log"
	"reflect"
	"testing"

	"github.com/spf13/cast"
	"github.com/theyakka/ystore"
	"github.com/theyakka/ystore/drivers/json"
)

func exampleStore() *ystore.Store {
	s := ystore.NewStore(ystore.WithCache(true, 5))
	_ = ystore.Set(s, "test.string", "hello")
	_ = ystore.Set(s, "test.int64", 122333)
	_ = ystore.Set(s, "test.bool.t", true)
	_ = ystore.Set(s, "test.bool.f", false)
	s.Debug()
	return s

}

func TestJSONDriver(t *testing.T) {
	s := ystore.NewStore(ystore.WithCache(true, 5))
	s.SetDriver(json.NewDriver())
	_ = s.Load("_testfiles/test.json")
}

func BenchmarkJSONDriver(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := ystore.NewStore(ystore.WithCache(true, 5))
		s.SetDriver(json.NewDriver())
		_ = s.Load("_testfiles/test.json")
	}
}

func TestMerge(t *testing.T) {
	s1 := ystore.NewStore()
	ystore.Set(s1, "test", 22)
	ystore.Set(s1, "test.this", 55)
	ystore.Set(s1, "test.that", 66)
	s2 := ystore.NewStore()
	ystore.Set(s2, "test.this", 99)
	ystore.Set(s2, "test.that.thing", 100)
	s3 := ystore.NewStore()
	ystore.Set(s3, "test.that.thing", "200")
	ystore.Set(s3, "test.that.thing.again", "monkey")
	ystore.Set(s3, "test.another", "turnip")

	ms, _ := ystore.Merge([]*ystore.Store{s1, s2, s3})
	log.Println("ms", "test.this", ms.Get("test.this").Value())
	log.Println("ms", "test.that", ms.Get("test.that").Value())
	log.Println("ms", "test.that.thing", ms.Get("test.that.thing").Value())
	log.Println("ms", "test.that.thing.again", ms.Get("test.that.thing.again").Value())
	log.Println("ms", "test.another", ms.Get("test.another").Value())
}

func TestSimpleGet(t *testing.T) {
	store := exampleStore()
	e := store.Get("test.bool.t")
	if e == nil {
		t.Fail()
		return
	}
	if e.Kind() != reflect.Bool {
		t.Fail()
		return
	}
	if cast.ToBool(e.Value()) != true {
		t.Fail()
		return
	}
}
