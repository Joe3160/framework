package cache

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/goravel/framework/contracts/cache"
	"github.com/goravel/framework/facades"

	"github.com/go-redis/redis/v8"
	"github.com/gookit/color"
)

type Redis struct {
	ctx    context.Context
	prefix string
	redis  *redis.Client
}

func NewRedis(ctx context.Context) cache.Store {
	connection := facades.Config.GetString("cache.stores." + facades.Config.GetString("cache.default") + ".connection")
	if connection == "" {
		connection = "default"
	}

	host := facades.Config.GetString("database.redis." + connection + ".host")
	if host == "" {
		return nil
	}

	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + facades.Config.GetString("database.redis."+connection+".port"),
		Password: facades.Config.GetString("database.redis." + connection + ".password"),
		DB:       facades.Config.GetInt("database.redis." + connection + ".database"),
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		color.Redln(fmt.Sprintf("[Cache] Init connection error, %s", err.Error()))

		return nil
	}

	return &Redis{
		ctx:    ctx,
		prefix: facades.Config.GetString("cache.prefix" + ":"),
		redis:  client,
	}
}

func (r *Redis) WithContext(ctx context.Context) cache.Store {
	return NewRedis(ctx)
}

//Get Retrieve an item from the cache by key.
func (r *Redis) Get(key string, def interface{}) interface{} {
	val, err := r.redis.Get(r.ctx, r.prefix+key).Result()
	if err != nil {
		switch s := def.(type) {
		case func() interface{}:
			return s()
		default:
			return def
		}
	}

	return val
}

func (r *Redis) GetBool(key string, def bool) bool {
	res := r.Get(key, def)
	if val, ok := res.(string); ok {
		return val == "1"
	}

	return res.(bool)
}

func (r *Redis) GetInt(key string, def int) int {
	res := r.Get(key, def)
	if val, ok := res.(string); ok {
		i, err := strconv.Atoi(val)
		if err != nil {
			return def
		}

		return i
	}

	return res.(int)
}

func (r *Redis) GetString(key string, def string) string {
	return r.Get(key, def).(string)
}

//Has Check an item exists in the cache.
func (r *Redis) Has(key string) bool {
	value, err := r.redis.Exists(r.ctx, r.prefix+key).Result()

	if err != nil || value == 0 {
		return false
	}

	return true
}

//Put Store an item in the cache for a given number of seconds.
func (r *Redis) Put(key string, value interface{}, seconds time.Duration) error {
	err := r.redis.Set(r.ctx, r.prefix+key, value, seconds).Err()
	if err != nil {
		return err
	}

	return nil
}

//Pull Retrieve an item from the cache and delete it.
func (r *Redis) Pull(key string, def interface{}) interface{} {
	val, err := r.redis.Get(r.ctx, r.prefix+key).Result()
	r.redis.Del(r.ctx, r.prefix+key)

	if err != nil {
		return def
	}

	return val
}

//Add Store an item in the cache if the key does not exist.
func (r *Redis) Add(key string, value interface{}, seconds time.Duration) bool {
	val, err := r.redis.SetNX(r.ctx, r.prefix+key, value, seconds).Result()
	if err != nil {
		return false
	}

	return val
}

//Remember Get an item from the cache, or execute the given Closure and store the result.
func (r *Redis) Remember(key string, ttl time.Duration, callback func() interface{}) (interface{}, error) {
	val := r.Get(key, nil)

	if val != nil {
		return val, nil
	}

	val = callback()

	if err := r.Put(key, val, ttl); err != nil {
		return nil, err
	}

	return val, nil
}

//RememberForever Get an item from the cache, or execute the given Closure and store the result forever.
func (r *Redis) RememberForever(key string, callback func() interface{}) (interface{}, error) {
	val := r.Get(key, nil)

	if val != nil {
		return val, nil
	}

	val = callback()

	if err := r.Put(key, val, 0); err != nil {
		return nil, err
	}

	return val, nil
}

//Forever Store an item in the cache indefinitely.
func (r *Redis) Forever(key string, value interface{}) bool {
	if err := r.Put(key, value, 0); err != nil {
		return false
	}

	return true
}

//Forget Remove an item from the cache.
func (r *Redis) Forget(key string) bool {
	_, err := r.redis.Del(r.ctx, r.prefix+key).Result()

	if err != nil {
		return false
	}

	return true
}

//Flush Remove all items from the cache.
func (r *Redis) Flush() bool {
	res, err := r.redis.FlushAll(r.ctx).Result()

	if err != nil || res != "OK" {
		return false
	}

	return true
}
