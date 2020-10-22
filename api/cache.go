package api

import (
	gocache "github.com/patrickmn/go-cache"
	"time"
)

type Cache interface {
	Put(string, []byte) error
	Get(string) ([]byte, bool, error)
	Delete(string) error
}

type GoCache struct {
	cache      *gocache.Cache
	expiration time.Duration
}

func NewGoCache(expiration time.Duration) *GoCache {
	return &GoCache{
		cache:      gocache.New(expiration, expiration),
		expiration: expiration,
	}
}

func (g *GoCache) Put(key string, data []byte) error {
	g.cache.Set(key, data, g.expiration)
	return nil
}

func (g *GoCache) Get(key string) ([]byte, bool, error) {
	value, present := g.cache.Get(key)
	if !present {
		return nil, false, nil
	}

	return value.([]byte), true, nil
}

func (g *GoCache) Delete(key string) error {
	g.cache.Delete(key)
	return nil
}
