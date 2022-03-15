package ystore

import "reflect"

func Cast[F any, T any](from F) (out T) {
	v := reflect.ValueOf(from)
	outType := reflect.TypeOf(out)
	if v.CanConvert(outType) {
		return v.Convert(outType).Interface().(T)
	}
	return reflect.Zero(outType).Interface().(T)
}

func CastD[F any, T any](from F, defaultValue T) (out T) {
	v := reflect.ValueOf(from)
	outType := reflect.TypeOf(out)
	if v.CanConvert(outType) {
		return v.Convert(outType).Interface().(T)
	}
	return defaultValue
}

func CastSlice[VT any](from any) []VT {
	a := reflect.ValueOf(from)
	v := a.Kind()
	if v != reflect.Slice {
		return nil
	}
	var out []VT
	for i := 0; i < a.Len(); i++ {
		out = append(out, a.Index(i).Interface().(VT))
	}
	return out
}
