package geeCache

import (
	"cache/geeCache/lru"
	"sync"
)

type cache struct {
	mutex sync.Mutex
	lru *lru.Cache
	maxByte int64
}

func (c *cache) get(key string)(value ByteView, ok bool)  {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.lru == nil {
		return               //todo 这里应该可以返回一个error
	}
	if value, ok := c.lru.Get(key); ok {
		return value.(ByteView), ok //强转
	}
	return
}

//从缓存中获取值，无法从源数据获取
func (c *cache) add(key string, value ByteView)  {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.lru == nil {
		return
	}
	c.lru.Add(key, value)
}


