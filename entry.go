// Created by Yakka (https://theyakka.com)
//
// Copyright (c) 2022 Yakka LLC.
// All rights reserved.
// See the LICENSE file for licensing details and requirements.

package ystore

import "reflect"

type Entry struct {
	store    *Store
	key      string
	value    reflect.Value
	parent   *Entry
	children map[string]*Entry
}
