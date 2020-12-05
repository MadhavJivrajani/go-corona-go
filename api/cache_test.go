package api

import (
	"io/ioutil"
	"path"
	"reflect"
	"testing"
	"time"
)

func TestGoCachePersistenceUnexistingFile(t *testing.T) {
	tempCacheFile, err := ioutil.TempDir("", "cache_*")
	if err != nil {
		t.Fatal(err)
	}

	cacheFile := path.Join(tempCacheFile, "cache")
	writer := NewGoCache(cacheFile,1 * time.Minute)
	reader := NewGoCache(cacheFile,1 * time.Minute)

	var (
		key   = "test"
		value = "test-value"
	)

	err = writer.Put(key, value)
	if err != nil {
		t.Fatal(err)
	}

	cached, present, err := reader.Get(key)
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

func TestGoCachePersistence(t *testing.T) {
	tempCacheFile, err := ioutil.TempFile("", "cache_*")
	if err != nil {
		t.Fatal(err)
	}

	writer := NewGoCache(tempCacheFile.Name(),1 * time.Minute)
	reader := NewGoCache(tempCacheFile.Name(),1 * time.Minute)

	var (
		key   = "test"
		value = "test-value"
	)

	err = writer.Put(key, value)
	if err != nil {
		t.Fatal(err)
	}

	cached, present, err := reader.Get(key)
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

func TestGoCachePutGet(t *testing.T) {
	tempCacheFile,err := ioutil.TempFile("", "cache_*")
	if err != nil{
		t.Fatal(err)
	}

	cache := NewGoCache(tempCacheFile.Name(),1 * time.Minute)

	var (
		key   = "test"
		value = "test-value"
	)

	err = cache.Put(key, value)
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
	tempCacheFile,err := ioutil.TempFile("", "cache_*")
	if err != nil{
		t.Fatal(err)
	}

	cache := NewGoCache(tempCacheFile.Name(),1 * time.Minute)

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
	tempCacheFile,err := ioutil.TempFile("", "cache_*")
	if err != nil{
		t.Fatal(err)
	}

	cache := NewGoCache(tempCacheFile.Name(),1 * time.Minute)

	var (
		key   = "test"
		value = "test-value"
	)

	err = cache.Put(key, value)
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

