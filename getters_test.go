package ystore

import (
	"testing"
)

var gettersStore *Store

func setupGetters() {
	gettersStore = NewStore()
	gettersStore.Set("some.nested.value", "turkey")
	gettersStore.Set("i.am.testing.this", "hello")
	gettersStore.Set("boolean", true)
	gettersStore.Set("numeric", 1234)
}

func TestGetStrings(t *testing.T) {
	testValue := gettersStore.GetString("some.nested.value")
	if testValue != "turkey" {
		t.Fail()
	}
	testValue2 := gettersStore.GetString("another.sub.string")
	if testValue2 != "" {
		t.Fail()
	}
	testValue3 := gettersStore.GetStringD("another.sub.string", "mastodon")
	if testValue3 != "mastodon" {
		t.Fail()
	}
	testValue4 := gettersStore.GetStringD("i.am.testing.this", "goodbye")
	if testValue4 != "hello" {
		t.Fail()
	}
}

func TestGetBool(t *testing.T) {
	testValue := gettersStore.GetBool("boolean")
	if testValue != true {
		t.Fail()
	}
	testValue2 := gettersStore.GetBool("doesnotexist")
	if testValue2 != false {
		t.Fail()
	}
	testValue3 := gettersStore.GetBoolD("alsodoesnotexist", true)
	if testValue3 != true {
		t.Fail()
	}
	testValue4 := gettersStore.GetBoolD("boolean", false)
	if testValue4 != true {
		t.Fail()
	}
}

func TestNonExistentSubStore(t *testing.T) {
	substore := gettersStore.Store("missing")
	if substore != nil {
		t.Fail()
	}
}
