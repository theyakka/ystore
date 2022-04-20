// Created by Yakka (https://theyakka.com)
//
// Copyright (c) 2022 Yakka LLC.
// All rights reserved.
// See the LICENSE file for licensing details and requirements.

package ystore

type Flag uint8

const (
	ParseArrays Flag = 1 << iota
	ParseObjects
)
