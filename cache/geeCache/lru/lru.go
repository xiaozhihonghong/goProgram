package lru

import (
	"container/list"
	"fmt"
)

// 回调函数的作用就是可以添加自己想要的函数，在执行删除节点的时候可以执行该函数，也可以是在其他地方。
type Cache struct {
	ll *list.List
	maxCap int64
	useCap int64
	cache map[string]*list.Element
	OnEvicted func(key string, value Value)
}

type Entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}

// 实例化缓存
func New(maxCap int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		ll: list.New(),
		maxCap: maxCap,
		useCap: 0,
		cache: make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// 查找节点
func (c *Cache) Get(key string) (value Value, ok bool)  {
	elem, ok := c.cache[key]
	fmt.Println(elem)
	if elem, ok := c.cache[key]; ok {
		c.ll.MoveToFront(elem)
		kv := elem.Value.(*Entry)
		return kv.value, ok
	}
	return
}

//删除节点
func (c *Cache) RemoveOldest()  {
	elem := c.ll.Back()
	if elem != nil {
		c.ll.Remove(elem)
		kv := elem.Value.(*Entry)
		delete(c.cache, kv.key)
		c.useCap = c.useCap - int64(len(kv.key)) - int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}
//新增和修改节点
func (c *Cache) Add(key string, value Value)  {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*Entry)
		c.useCap += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushFront(&Entry{key: key, value: value})
		c.cache[key] = ele
		c.useCap += int64(len(key)) + int64(value.Len())
	}
	for c.useCap != 0 && c.useCap > c.maxCap {
		c.RemoveOldest()
	}
}