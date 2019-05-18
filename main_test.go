package ystore

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setupGetters()
	os.Exit(m.Run())
}
