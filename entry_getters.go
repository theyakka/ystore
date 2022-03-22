package ystore

import (
	"reflect"
	"strings"
)

func (e *Entry) Has(keyPath string) bool {
	e.store.mutex.RLock()
	entry := FindEntry(e.store, e.children, strings.Split(keyPath, "."))
	e.store.mutex.RUnlock()
	return entry != nil
}

func (e *Entry) Key() string {
	return e.key
}

func (e *Entry) KeyPath() string {
	var segments []string
	current := e
	for current != nil {
		segments = append(segments, current.key)
		current = current.parent
	}
	return strings.Join(segments, ".")
}

func (e *Entry) Parent() *Entry {
	return e.parent
}

func (e *Entry) Kind() reflect.Kind {
	return e.value.Kind()
}

func (e *Entry) IsValid() bool {
	if e == nil || !e.value.IsValid() {
		return false
	}
	switch e.value.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		if e.value.IsNil() {
			return false
		}
	}
	return true
}

func (e *Entry) Get(keyPath string) *Entry {
	e.store.mutex.RLock()
	entry := e.store.cache[e.KeyPath()]
	if entry != nil {
		e.store.mutex.RUnlock()
		return entry
	}
	entry = FindEntry(e.store, e.children, strings.Split(keyPath, "."))
	e.store.mutex.RUnlock()
	return entry
}

func (e *Entry) Value() any {
	if e != nil && e.value.IsValid() {
		return e.value.Interface()
	}
	return nil
}

func (e *Entry) ValueD(defaultValue any) any {
	if !e.IsValid() {
		return defaultValue
	}
	return e.value
}
