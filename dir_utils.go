package ystore

import "path/filepath"

func BaseDir(filename string) string {
	dir := filepath.Dir(filename)
	if dir != "" {
		return filepath.Base(dir)
	}
	return ""
}
