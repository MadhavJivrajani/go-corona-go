package api

import (
	"encoding/json"
	gocache "github.com/patrickmn/go-cache"
	"io/ioutil"
	"os"
	"time"
)

type Cache interface {
	Put(string, string) error
	Get(string) (string, bool, error)
	Delete(string) error
}

type GoCache struct {
	cache      *gocache.Cache
	expiration time.Duration
	path       string
	needSync   bool
}

func NewGoCache(path string, expiration time.Duration) *GoCache {
	return &GoCache{
		cache:      gocache.New(expiration, expiration),
		expiration: expiration,
		path:       path,
		needSync:   true,
	}
}

func (g *GoCache) Put(key string, data string) error {
	if err := g.syncIfNeeded(); err != nil {
		return err
	}

	g.cache.Set(key, data, g.expiration)
	return g.syncToFile()
}

func (g *GoCache) Get(key string) (string, bool, error) {
	if err := g.syncIfNeeded(); err != nil {
		return "", false, err
	}

	value, present := g.cache.Get(key)
	if !present {
		return "", false, nil
	}

	return value.(string), true, nil
}

func (g *GoCache) Delete(key string) error {
	if err := g.syncIfNeeded(); err != nil {
		return err
	}

	g.cache.Delete(key)
	return g.syncToFile()
}

func (g *GoCache) syncIfNeeded() error {
	if g.needSync {
		if err := g.syncFromFile(); err != nil {
			return err
		}
		g.needSync = false
	}
	return nil
}

func (g *GoCache) syncFromFile() error {
	content, err := ioutil.ReadFile(g.path)

	if os.IsNotExist(err) {
		return nil
	}

	if err != nil {
		return err
	}

	if len(content) == 0 {
		return nil
	}

	var state map[string]gocache.Item
	if err := json.Unmarshal(content, &state); err != nil {
		return err
	}

	g.cache = gocache.NewFrom(g.expiration, g.expiration, state)
	return nil
}

func (g *GoCache) syncToFile() error {
	items := g.cache.Items()
	serialized, err := json.Marshal(items)
	if err != nil {
		return err
	}

	fp, err := os.Create(g.path)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write(serialized)
	if err != nil {
		return nil
	}

	return nil
}
