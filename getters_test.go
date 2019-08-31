package ystore_test

import (
	"log"
	"testing"

	"github.com/theyakka/ystore"
)

var gettersStore *ystore.Store

const testInt = 1234
const testFloat = 98.7654

var testIntSlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
var testStringSlice = []string{
	"item 1",
	"item 2",
	"item 3",
}
var testMapValue = map[string]interface{}{
	"key1": 1234,
	"key2": "orange",
	"key3": 67.12,
}

func setupGetters() {
	gettersStore = ystore.NewStore()
	gettersStore.Set("some.nested.value", "turkey")
	gettersStore.Set("i.am.testing.this", "hello")
	gettersStore.Set("stringSlice", testStringSlice)
	gettersStore.Set("boolean", true)
	gettersStore.Set("int", testInt)
	gettersStore.Set("intSlice", testIntSlice)
	gettersStore.Set("float", testFloat)
	gettersStore.Set("map", testMapValue)
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

func TestStringValues(t *testing.T) {
	stringSlice := gettersStore.GetStringSlice("stringSlice")
	if !stringSliceEquals(stringSlice, testStringSlice) {
		t.Fail()
	}
	defaultSlice := []string{
		"banana", "apple",
	}
	goodStringSlice := gettersStore.GetStringSliceD("stringSlice", defaultSlice)
	if !stringSliceEquals(goodStringSlice, testStringSlice) {
		t.Fail()
	}
	badStringSlice := gettersStore.GetStringSliceD("nonExistStringSlice", defaultSlice)
	if !stringSliceEquals(badStringSlice, defaultSlice) {
		t.Fail()
	}
}

func TestIntValues(t *testing.T) {
	// integer values
	intVal := gettersStore.GetInt("int")
	if intVal != testInt {
		t.Fail()
	}
	goodIntVal := gettersStore.GetIntD("int", 11)
	if goodIntVal != testInt {
		t.Fail()
	}
	badIntVal := gettersStore.GetIntD("nonExistentInt", 50)
	if badIntVal != 50 {
		t.Fail()
	}
	defaultSlice := []int{
		234, 123, 567,
	}
	intSlice := gettersStore.GetIntSlice("intSlice")
	if !intSliceEquals(intSlice, testIntSlice) {
		t.Fail()
	}
	goodIntSlice := gettersStore.GetIntSliceD("intSlice", defaultSlice)
	if !intSliceEquals(goodIntSlice, testIntSlice) {
		t.Fail()
	}
	badIntSlice := gettersStore.GetIntSliceD("nonExistIntSlice", defaultSlice)
	if !intSliceEquals(badIntSlice, defaultSlice) {
		t.Fail()
	}
}

func TestFloatValues(t *testing.T) {
	// float values
	floatVal := gettersStore.GetFloat("float")
	if floatVal != testFloat {
		t.Fail()
	}
	goodFloatVal := gettersStore.GetFloatD("float", 11.32)
	if goodFloatVal != testFloat {
		t.Fail()
	}
	badFloatVal := gettersStore.GetFloatD("nonExistentFloat", 22.67)
	if badFloatVal != 22.67 {
		t.Fail()
	}
}

func TestMap(t *testing.T) {
	mapVal := gettersStore.GetMap("map")
	if !mapEquals(mapVal, testMapValue) {
		log.Println("fail1")
		t.Fail()
	}
	defaultMap := map[string]interface{}{
		"def1": "hello",
		"def2": 1234,
	}
	goodMapVal := gettersStore.GetMapD("map", defaultMap)
	if !mapEquals(goodMapVal, testMapValue) {
		log.Println("fail2")
		t.Fail()
	}
	badMapVal := gettersStore.GetMapD("nonExistMap", defaultMap)
	if !mapEquals(badMapVal, defaultMap) {
		log.Println("fail3")
		t.Fail()
	}
}

func intSliceEquals(slice1 []int, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for idx, val := range slice1 {
		if slice2[idx] != val {
			return false
		}
	}
	return true
}

func stringSliceEquals(slice1 []string, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for idx, val := range slice1 {
		if slice2[idx] != val {
			return false
		}
	}
	return true
}

func mapEquals(map1 map[string]interface{}, map2 map[string]interface{}) bool {
	if len(map1) != len(map2) {
		return false
	}
	for key, val := range map1 {
		if map2[key] != val {
			return false
		}
	}
	return true
}
