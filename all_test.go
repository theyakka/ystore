// Created by Yakka (https://theyakka.com)
//
// Copyright (c) 2022 Yakka LLC.
// All rights reserved.
// See the LICENSE file for licensing details and requirements.

package ystore_test

import (
	"fmt"
	"github.com/theyakka/ystore"
	"github.com/theyakka/ystore/drivers/json_driver"
	"testing"
)

func exampleStore() *ystore.Store {
	s := ystore.NewStore(ystore.WithCache(true, 5))
	_ = ystore.Set(s, "test.string", "hello")
	_ = ystore.Set(s, "test.int64", 30193214032540302392923i)
	_ = ystore.Set(s, "test.bool.t", true)
	_ = ystore.Set(s, "test.bool.f", false)
	return s

}

func TestJSONDriver(t *testing.T) {
	s := ystore.NewStore(ystore.WithCache(true, 5))
	s.SetDriver(json_driver.NewDriver())
	_ = s.Load("_testfiles/test.json_driver")
	e := s.Get("test.nested.array")
	fmt.Println(e.RawValue())
}

func TestSimpleGet(t *testing.T) {
	store := exampleStore()
	e := store.Get("test.bool.t")
	if e == nil {
		t.Fail()
		return
	}
	if e.BoolValue() == false {
		t.Fail()
		return
	}
}

func TestNestedGet(t *testing.T) {
	store := ystore.NewStore()
	_ = ystore.Set(store, "hello.this.thing", 55)
	_ = ystore.Set(store, "hello.this.other.thing", 72)
	_ = ystore.Set(store, "hello.this.that", 22)

	e1 := ystore.Get(store, "hello.this")
	e2 := e1.Get("other.thing")
	if e2 == nil {
		t.Fail()
		return
	}
	if e2.IntValue() != 72 {
		t.Fail()
		return
	}

	e3 := e1.Get("other").Get("thing")
	if e3 == nil {
		t.Fail()
		return
	}
	if e3.IntValue() != 72 {
		t.Fail()
		return
	}
}

func TestSimpleSet(t *testing.T) {
	store := exampleStore()
	e := store.Get("test")
	if e == nil {
		t.Fail()
		return
	}
	_ = e.Set("nested.thing", "banana")

	e = store.Get("test.nested.thing")
	if e == nil {
		t.Fail()
		return
	}
	if e.StringValue() != "banana" {
		t.Fail()
		return
	}

	_ = e.Set("hello", "turnips are cool")

	se := ystore.Get(store, "test.nested.thing.hello")
	if se == nil {
		t.Fail()
		return
	}
	if se.StringValue() != "turnips are cool" {
		t.Fail()
		return
	}
}
