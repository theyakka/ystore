package ystore

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestInvalidDataDir(t *testing.T) {
	baseDir, dirErr := filepath.Abs("./_doesnotexist")
	if dirErr != nil {
		t.Error(dirErr)
	}
	store := StoreWithDir(baseDir)
	storeErr := store.ReadAll()
	if storeErr != nil {
		return	// should fail
	}
	t.Error("This should have failed because the directory doesn't exist")
}

func TestInvalidFile(t *testing.T) {
	baseDir, dirErr := filepath.Abs("./_testdata/first.toml")
	if dirErr != nil {
		t.Error(dirErr)
	}
	store := StoreWithDir(baseDir)
	storeErr := store.ReadAll()
	if storeErr != nil {
		return	// should fail
	}
	t.Error("This should have failed because the directory doesn't exist")
}

func TestParsingDataDir(t *testing.T) {
	baseDir, dirErr := filepath.Abs("./_testdata")
	if dirErr != nil {
		t.Error(dirErr)
		return
	}
	store := StoreWithDir(baseDir)
	storeErr := store.ReadAll()
	if storeErr != nil {
		t.Error(storeErr)
		return
	}

	fmt.Println(store.AllValues())
	fmt.Println(store.Get("categories.firstcategory.name"))
}
