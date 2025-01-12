package cache

import (
	"context"
	"time"
)

//go:generate mockery --name=Store
type Store interface {
	WithContext(ctx context.Context) Store
	//Get Retrieve an item from the cache by key.
	Get(key string, def interface{}) interface{}
	GetBool(key string, def bool) bool
	GetInt(key string, def int) int
	GetString(key string, def string) string
	//Has Check an item exists in the cache.
	Has(key string) bool
	//Put Store an item in the cache for a given number of seconds.
	Put(key string, value interface{}, sec time.Duration) error
	//Pull Retrieve an item from the cache and delete it.
	Pull(key string, def interface{}) interface{}
	//Add Store an item in the cache if the key does not exist.
	Add(key string, value interface{}, sec time.Duration) bool
	//Remember Get an item from the cache, or execute the given Closure and store the result.
	Remember(key string, ttl time.Duration, callback func() interface{}) (interface{}, error)
	//RememberForever Get an item from the cache, or execute the given Closure and store the result forever.
	RememberForever(key string, callback func() interface{}) (interface{}, error)
	//Forever Store an item in the cache indefinitely.
	Forever(key string, value interface{}) bool
	//Forget Remove an item from the cache.
	Forget(key string) bool
	//Flush Remove all items from the cache.
	Flush() bool
}
