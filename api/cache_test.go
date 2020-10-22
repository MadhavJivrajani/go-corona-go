package api

import (
	"reflect"
	"testing"
	"time"
)

func TestGoCachePutGet(t *testing.T) {
	cache := NewGoCache(1 * time.Minute)

	var (
		key   = "test"
		value = []byte("test-value")
	)

	err := cache.Put(key, value)
	if err != nil {
		t.Fatal(err)
	}

	cached, present, err := cache.Get(key)
	if err != nil {
		t.Fatal(err)
	}

	if !present {
		t.Fatalf("key %s must be present", key)
	}

	if !reflect.DeepEqual(value, cached) {
		t.Fatal("cached values is not equals to written value")
	}
}


func TestGoCacheGetNonExistentKey(t *testing.T) {
	cache := NewGoCache(1 * time.Minute)

	var (
		key   = "test"
	)

	_, present, err := cache.Get(key)
	if err != nil {
		t.Fatal(err)
	}

	if present {
		t.Fatalf("key %s must be not present", key)
	}


}



func TestGoCachePutDeleteGet(t *testing.T) {
	cache := NewGoCache(1 * time.Minute)

	var (
		key   = "test"
		value = []byte("test-value")
	)

	err := cache.Put(key, value)
	if err != nil {
		t.Fatal(err)
	}

	err = cache.Delete(key)
	if err != nil {
		t.Fatal(err)
	}

	_, present, err := cache.Get(key)
	if err != nil {
		t.Fatal(err)
	}

	if present {
		t.Fatalf("key %s must be not present", key)
	}
}

