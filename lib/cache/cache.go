package cache

import (
	"encoding/json"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	_pool *redis.Pool
)

// Open open
func Open(pool *redis.Pool) {
	_pool = pool
}

// Set set bytes
func Set(key string, val interface{}, ttl time.Duration) error {
	buf, err := json.Marshal(val)
	if err != nil {
		return err
	}

	c := _pool.Get()
	defer c.Close()
	_, err = c.Do("SET", cacheKey(key), buf, "EX", int(ttl/time.Second))
	return err
}

// Get get bytes
func Get(key string, val interface{}) error {

	c := _pool.Get()
	defer c.Close()
	buf, err := redis.Bytes(c.Do("GET", cacheKey(key)))
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, val)
}

// Status status
func Status() (string, error) {
	c := _pool.Get()
	defer c.Close()
	return redis.String(c.Do("INFO"))
}

//Flush clear cache items
func Flush() error {
	c := _pool.Get()
	defer c.Close()
	keys, err := redis.Values(c.Do("KEYS", cacheKey("*")))
	if err == nil && len(keys) > 0 {
		_, err = c.Do("DEL", keys...)
	}
	return err
}

func cacheKey(k string) string {
	return "cache://" + k
}
