// Created by Yakka (https://theyakka.com)
//
// Copyright (c) 2020 Yakka LLC.
// All rights reserved.
// See the LICENSE file for licensing details and requirements.

package ystore

import "path/filepath"

func BaseDir(filename string) string {
	return filepath.Base(filepath.Dir(filename))
}
