// Package cache
// 缓存模块
package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type Cache struct {
	C *cache.Cache
}

func SetUp() *Cache {
	var newCache Cache
	newCache.C = cache.New(5*time.Minute, 10*time.Minute)
	return &newCache
}

// Set 存储
func (c *Cache) Set(key string, data interface{}, date time.Duration) {
	c.C.Set(key, data, date)
}

// Get 获取
func (c *Cache) Get(key string) (d interface{}, found bool) {
	return c.C.Get(key)
}

// Delete 删除缓存
func (c *Cache) Delete(key string) {
	c.C.Delete(key)
}

// Flush 刷新，删除所有缓存。
func (c *Cache) Flush() {
	c.C.Flush()
}

// SetBytes 把byte切片数据放进缓存中
func (c *Cache) SetBytes(key string, data []byte, date time.Duration) {
	c.C.Set(key, data, date)
}

// GetBytes 获取[]byte类型的缓存
func (c *Cache) GetBytes(key string) ([]byte, bool) {
	if x, found := c.C.Get(key); found {
		bytes, ok := x.([]byte)
		if !ok {
			return nil, false
		}
		return bytes, true
	}
	return nil, false
}

// Exists 判断key是否存在
func (c *Cache) Exists(key string) bool {
	_, isExist := c.C.Get(key)
	return isExist
}
