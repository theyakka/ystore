package ystore

import "path/filepath"

func BaseDir(filename string) string {
	return filepath.Base(filepath.Dir(filename))
}
