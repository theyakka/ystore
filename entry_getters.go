package ystore

import (
	"github.com/spf13/cast"
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
	if e == nil || e.value.IsNil() || e.value.IsZero() {
		return defaultValue
	}
	return e.value
}

func (e *Entry) StringValue() string {
	return cast.ToString(e.RawValue())
}

func (e *Entry) StringValueD(defaultValue string) string {
	return cast.ToString(e.RawValueD(defaultValue))
}

func (e *Entry) BoolValue() bool {
	return cast.ToBool(e.RawValue())
}

func (e *Entry) BoolValueD(defaultValue bool) bool {
	return cast.ToBool(e.RawValueD(defaultValue))
}

func (e *Entry) IntValue() int {
	return cast.ToInt(e.RawValue())
}

func (e *Entry) IntValueD(defaultValue int) int {
	return cast.ToInt(e.RawValueD(defaultValue))
}

func (e *Entry) FloatValue() float64 {
	return cast.ToFloat64(e.RawValue())
}

func (e *Entry) FloatValueD(defaultValue float64) float64 {
	return cast.ToFloat64(e.RawValueD(defaultValue))
}
