package ystore

import "reflect"

func Cast[F any, T any](from F) T {
	return reflect.ValueOf(from).Interface().(T)
}

func CastSlice[F any, VT any](from F) []VT {
	a := reflect.ValueOf(from).Interface()

}
