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
	storeErr := store.ReadDir(baseDir)
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
	storeErr := store.ReadDir(baseDir)
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
	fmt.Println(store.AllValues())
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
	fmt.Println(store.AllValues())

}
