package ystore_test

import (
	"fmt"
	"testing"

	"github.com/theyakka/ystore"
)

func TestProcessing(t *testing.T) {
	store := ystore.NewStoreFromMapWithSubs(map[string]interface{}{
		"first":  "hello %s",
		"second": 1234,
		"third": map[string]interface{}{
			"threeone": "testing %s",
		},
	})
	newStore := store.Process(func(k string, v interface{}) interface{} {
		switch v.(type) {
		case string:
			return fmt.Sprintf(v.(string), k)
		default:
			return v
		}
	})
	if newStore.Len() != store.Len() {
		t.Error("store lengths don't match")
		return
	}
	strVal := newStore.GetString("third.threeone")
	if strVal != "testing threeone" {
		t.Error("expected string value is incorrect")
		return
	}
	intVal := newStore.GetInt("second")
	if intVal != 1234 {
		t.Error("expected int value is incorrect")
		return
	}
}
