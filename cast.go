package ystore

import "reflect"

func Cast[T any](value reflect.Value) (out T) {
	outType := reflect.TypeOf(out)
	if value.CanConvert(outType) {
		return value.Convert(outType).Interface().(T)
	}
	return reflect.Zero(outType).Interface().(T)
}

func CastD[T any](value reflect.Value, defaultValue T) (out T) {
	outType := reflect.TypeOf(out)
	if value.CanConvert(outType) {
		return value.Convert(outType).Interface().(T)
	}
	return defaultValue
}

func CastSlice[VT any](value reflect.Value) []VT {
	v := value.Kind()
	if v != reflect.Slice {
		return nil
	}
	var out []VT
	for i := 0; i < value.Len(); i++ {
		out = append(out, value.Index(i).Interface().(VT))
	}
	return out
}
