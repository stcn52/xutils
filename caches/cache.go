package caches

import "sync"

// 缓存集成器，用户让别的对象拥有使用缓存的能力
type CacheFactory struct {
	factory *Factory
	locker  sync.Once
}

// 取得缓存管理器
func (c *CacheFactory) Cache() *Factory {
	c.locker.Do(func() {
		c.factory = NewFactory()
	})
	return c.factory
}
