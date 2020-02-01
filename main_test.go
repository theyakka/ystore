// Created by Yakka (https://theyakka.com)
//
// Copyright (c) 2020 Yakka LLC.
// All rights reserved.
// See the LICENSE file for licensing details and requirements.

package ystore_test

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setupGetters()
	os.Exit(m.Run())
}
