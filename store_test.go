// Created by Yakka (https://theyakka.com)
//
// Copyright (c) 2020 Yakka LLC.
// All rights reserved.
// See the LICENSE file for licensing details and requirements.

package ystore_test

import (
	"errors"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/theyakka/ystore"
)

func TestStoreOrEmpty(t *testing.T) {
	store := ystore.NewStore()
	store.Set("substore", map[string]interface{}{
		"first":  "hello",
		"second": 1234,
		"third":  12.34,
		"fourth": []int{1, 2, 3, 4},
	})
	store.Set("simple", 55)
	substore := store.StoreFromMapOrEmpty("nonexisting")
	if substore == nil {
		t.Error("substore should not be nil")
		return
	}
	substore = store.StoreFromMap("somethingnothere")
	if substore != nil {
		t.Error("substore should be nil")
		return
	}
	substore = store.StoreFromMapOrEmpty("simple")
	if substore == nil {
		t.Error("substore should not be nil")
		return
	}
	if substore.Len() > 0 {
		t.Error("store should be empty")
		return
	}
}

func TestPassingSimpleMap(t *testing.T) {
	testMap := map[string]interface{}{
		"first":  "hello",
		"second": 1234,
		"third":  12.34,
		"fourth": []int{1, 2, 3, 4},
	}
	testStore := ystore.NewStoreFromMap(testMap)
	if testStore.Len() != len(testMap) {
		t.Fail()
	}
}

func TestPassingNestedMap(t *testing.T) {
	nestedMap2 := map[string]interface{}{
		"color1": "red",
		"color2": "green",
		"color3": "blue",
		"color4": "purple",
	}
	nestedMap1 := map[string]interface{}{
		"nested1": "nestedhello",
		"nested2": 4321,
		"colors":  nestedMap2,
	}
	testMap := map[string]interface{}{
		"first":  "hello",
		"second": nestedMap1,
	}
	testStore := ystore.NewStoreFromMap(testMap)
	if testStore.Len() != len(testMap) {
		t.Fail()
	}
	substore1 := testStore.StoreFromMapOrEmpty("second")
	if substore1.Len() != len(nestedMap1) {
		t.Fail()
	}
	substore2 := testStore.StoreFromMapOrEmpty("second.colors")
	if substore2.Len() != len(nestedMap2) {
		t.Fail()
	}
}

func TestFailWhenPassingNonExistentDir(t *testing.T) {
	baseDir, dirErr := filepath.Abs("./_doesnotexist")
	if dirErr != nil {
		t.Error(dirErr)
	}
	store := ystore.NewStore()
	storeErr := store.ReadDir(baseDir)
	if storeErr != nil {
		return // should fail
	}
	t.Error("This should have failed because the directory doesn't exist")
}

func TestFailWhenPassingFileToReadAll(t *testing.T) {
	baseDir, dirErr := filepath.Abs("./_testdata/first.toml")
	if dirErr != nil {
		t.Error(dirErr)
	}
	store := ystore.NewStore()
	storeErr := store.ReadDir(baseDir)
	if storeErr != nil {
		return // should fail
	}
	t.Error("This should have failed because the directory doesn't exist")
}

func TestParseAllInDirectory(t *testing.T) {
	baseDir, dirErr := filepath.Abs("./_testdata")
	if dirErr != nil {
		t.Error(dirErr)
		return
	}
	store := ystore.NewStore()
	storeErr := store.ReadDir(baseDir)
	if storeErr != nil {
		t.Error(storeErr)
		return
	}
}

func TestParseSingleFile(t *testing.T) {
	filename, dirErr := filepath.Abs("./_testdata/first.toml")
	if dirErr != nil {
		t.Error(dirErr)
		return
	}
	store := ystore.NewStore()
	if storeErr := store.ReadFile(filename); storeErr != nil {
		t.Error(storeErr)
		return
	}
}

func TestParseSingleBadYamlFile(t *testing.T) {
	filename, dirErr := filepath.Abs("./_badtestdata/bad.yaml")
	if dirErr != nil {
		t.Error(dirErr)
		return
	}
	store := ystore.NewStore()
	if storeErr := store.ReadFile(filename); storeErr == nil {
		t.Error(errors.New("file is bad and should have failed"))
		return
	}
}

func TestParseNonExistingFile(t *testing.T) {
	store := ystore.NewStore()
	if storeErr := store.ReadFile("./_badtestdata/nonexistent.yaml"); storeErr == nil {
		t.Error(errors.New("file is non-existent and should have failed"))
		return
	}
}

func TestParseMultipleFiles(t *testing.T) {
	var originalFilenames = []string{
		"./_testdata/first.toml",
		"./_testdata/third.json",
	}
	var filenames []string
	for _, filename := range originalFilenames {
		absFilename, dirErr := filepath.Abs(filename)
		if dirErr != nil {
			t.Error(dirErr)
			return
		}
		filenames = append(filenames, absFilename)
	}
	store := ystore.NewStore()
	if storeErr := store.ReadFiles(filenames...); storeErr != nil {
		t.Error(storeErr)
		return
	}
}

func TestStoreMerging(t *testing.T) {
	store1 := ystore.NewStoreFromMap(map[string]interface{}{
		"item1": "item 1",
		"item2": 567,
		"item3": []string{"string 1", "string 2"},
	})
	store2 := ystore.NewStoreFromMap(map[string]interface{}{
		"item4": 234.567,
		"item5": "item 5",
	})
	mergedStore := ystore.MergeStores(*store1, *store2)
	if mergedStore.Len() != 5 {
		t.Fail()
	}
}

func ExampleStore() {
	store := ystore.NewStore()
	store.Set("color", "red")
	store.Set("length", 100)
	fmt.Printf("The item is %s and the length is %d.\n", store.GetString("color"), store.GetInt("length"))
	// Output: The item is red and the length is 100.
}

func ExampleStore_map() {
	store := ystore.NewStoreFromMap(map[string]interface{}{
		"color":  "green",
		"length": 80,
	})
	fmt.Printf("The item is %s and the length is %d.\n", store.GetString("color"), store.GetInt("length"))
	// Output: The item is green and the length is 80.
}

func ExampleStore_file() {
	filename, fileErr := filepath.Abs("./_testdata/first.toml")
	if fileErr != nil {
		// handle the path error
		return
	}
	store := ystore.NewStore()
	if readErr := store.ReadFile(filename); readErr != nil {
		// handle store load error
		return
	}
	name := store.GetString("firstdata.name")
	numbers := store.GetIntSlice("firstdata.numbers")
	fmt.Printf("The item name is '%s' and the numbers slice has %d element(s).\n", name, len(numbers))
	// Output: The item name is 'First Item' and the numbers slice has 5 element(s).
}
