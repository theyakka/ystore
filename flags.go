package ystore

type Flag uint8

const (
	ParseArrays Flag = 1 << iota
	ParseObjects
)
