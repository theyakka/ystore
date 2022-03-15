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

func (e *Entry) RawValue() any {
	if e != nil && e.value.IsValid() {
		return e.value.Interface()
	}
	return nil
}

func (e *Entry) RawValueD(defaultValue any) any {
	if e == nil {
		return defaultValue
	}
	switch reflect.TypeOf(e.value).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		if e.value.IsNil() {
			return defaultValue
		}
	}
	return e.value
}

func (e *Entry) StringValue() string {
	return Cast[any, string](e.RawValue())
}

func (e *Entry) StringValueD(defaultValue string) string {
	return Cast[any, string](e.RawValueD(defaultValue))
}

func (e *Entry) BoolValue() bool {
	return Cast[any, bool](e.RawValue())
}

func (e *Entry) BoolValueD(defaultValue bool) bool {
	return Cast[any, bool](e.RawValueD(defaultValue))
}

func (e *Entry) IntValue() int {
	return Cast[any, int](e.RawValue())
}

func (e *Entry) IntValueD(defaultValue int) int {
	return CastD[any, int](e.RawValue(), defaultValue)
}

func (e *Entry) FloatValue() float64 {
	return Cast[any, float64](e.RawValue())
}

func (e *Entry) FloatValueD(defaultValue float64) float64 {
	return Cast[any, float64](e.RawValueD(defaultValue))
}

func (e *Entry) SliceValue() []any {
	return CastSlice[any](e.RawValue())
}
