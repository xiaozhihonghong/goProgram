package cache

import (
	"fmt"
	"reflect"
	"testing"
)

// 匹配value的接口类型
type String string

//必须实现接口中的Len()
func (d String) Len() int  {
	return len(d)
}


// 注意 文件名称必须是xxx_test的形式，否则无法运行
func TestGet(t *testing.T)  {
	lru := New(int64(10000), nil)
	lru.Add("key1", String("123"))
	fmt.Println(lru)
	if value, ok := lru.Get("key1"); !ok || string(value.(String)) != "123" {
		fmt.Println(value)
		t.Fatalf("cache hit key1=123 failed")
	}
	if _, ok := lru.Get("key2"); ok {
		t.Fatalf("cache miss key2 failed")
	}
}

func TestRemoveoldest(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "k3"
	v1, v2, v3 := "value1", "value2", "v3"
	caps := len(k1 + k2 + v1 + v2)
	lru := New(int64(caps), nil)
	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	lru.Add(k3, String(v3))

	if _, ok := lru.Get("key1"); ok || len(lru.cache) != 2 {
		t.Fatalf("Removeoldest key1 failed")
	}
}

func TestOnEvicted(t *testing.T) {
	keys := make([]string, 0)
	callback := func(key string, value Value) {
		keys = append(keys, key)
	}
	lru := New(int64(10), callback)
	lru.Add("key1", String("123456"))
	lru.Add("k2", String("k2"))
	lru.Add("k3", String("k3"))
	lru.Add("k4", String("k4"))

	expect := []string{"key1", "k2"}

	if !reflect.DeepEqual(expect, keys) {
		t.Fatalf("Call OnEvicted failed, expect keys equals to %s", expect)
	}
}
