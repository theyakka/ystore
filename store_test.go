package ystore

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestFailWhenPassingNonExistentDir(t *testing.T) {
	baseDir, dirErr := filepath.Abs("./_doesnotexist")
	if dirErr != nil {
		t.Error(dirErr)
	}
	store := NewStore()
	storeErr := store.ReadAll(baseDir)
	if storeErr != nil {
		return	// should fail
	}
	t.Error("This should have failed because the directory doesn't exist")
}

func TestFailWhenPassingFileToReadAll(t *testing.T) {
	baseDir, dirErr := filepath.Abs("./_testdata/first.toml")
	if dirErr != nil {
		t.Error(dirErr)
	}
	store := NewStore()
	storeErr := store.ReadAll(baseDir)
	if storeErr != nil {
		return	// should fail
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
	storeErr := store.ReadAll(baseDir)
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
	fmt.Println(store.AllValues())
}
