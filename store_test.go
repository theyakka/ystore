package ystore

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestPassingSimpleMap(t *testing.T) {
	testMap := map[string]interface{}{
		"first":  "hello",
		"second": 1234,
		"third":  12.34,
		"fourth": []int{1, 2, 3, 4},
	}
	testStore := NewStoreFromMap(testMap)
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
	testStore := NewStoreFromMap(testMap)
	if testStore.Len() != len(testMap) {
		t.Fail()
	}
	substore1 := testStore.Store("second")
	if substore1.Len() != len(nestedMap1) {
		t.Fail()
	}
	substore2 := testStore.Store("second.colors")
	if substore2.Len() != len(nestedMap2) {
		t.Fail()
	}
}

func TestFailWhenPassingNonExistentDir(t *testing.T) {
	baseDir, dirErr := filepath.Abs("./_doesnotexist")
	if dirErr != nil {
		t.Error(dirErr)
	}
	store := NewStore()
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
	store := NewStore()
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
	store := NewStore()
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
	store := NewStore()
	if storeErr := store.ReadFile(filename); storeErr != nil {
		t.Error(storeErr)
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
	store := NewStore()
	if storeErr := store.ReadFiles(filenames...); storeErr != nil {
		t.Error(storeErr)
		return
	}
}

func TestStoreMerging(t *testing.T) {
	store1 := NewStoreFromMap(map[string]interface{}{
		"item1": "item 1",
		"item2": 567,
		"item3": []string{"string 1", "string 2"},
	})
	store2 := NewStoreFromMap(map[string]interface{}{
		"item4": 234.567,
		"item5": "item 5",
	})
	mergedStore := MergeStores(store1, store2)
	if mergedStore.Len() != 5 {
		t.Fail()
	}
}

func ExampleStore() {
	store := NewStore()
	store.Set("color", "red")
	store.Set("length", 100)
	fmt.Printf("The item is %s and the length is %d.\n", store.GetString("color"), store.GetInt("length"))
	// Output: The item is red and the length is 100.
}

func ExampleStore_newmap() {
	store := NewStoreFromMap(map[string]interface{}{
		"color":  "green",
		"length": 80,
	})
	fmt.Printf("The item is %s and the length is %d.\n", store.GetString("color"), store.GetInt("length"))
	// Output: The item is green and the length is 80.
}
