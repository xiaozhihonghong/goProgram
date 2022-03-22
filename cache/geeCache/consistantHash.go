package geeCache

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func(data []byte)uint32

type Map struct {
	hash Hash //哈希算法
	replicas int //每个真实节点的虚拟节点倍数
	keys []int  //哈希环
	hashMap map[int]string //虚拟节点和真实节点的映射
}

func NewMap(replicas int, h Hash) *Map {
	m := &Map{
		hash: h,
		replicas: replicas,
		hashMap: make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE    //一致性hash算法
	}
	return m
}

//添加节点，解决资源倾斜问题，扩容
func (m *Map) Add(keys ...string)  {
	for _, key := range keys {
		for i:=0;i<m.replicas;i++ {
			//计算虚拟节点的hash值
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			//添加到hash环中
			m.keys = append(m.keys, hash)
			//建立虚拟节点到真实节点的映射
			m.hashMap[hash] = key
		}
	}
}

//寻找节点
func (m *Map) Get(key string) string {
	// 虚拟节点判空
	// 计算虚拟节点
	//映射找到真实节点
	if len(m.keys) == 0 {
		return ""
	}

	h := int(m.hash([]byte(key)))

	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= h
	})
	return m.hashMap[m.keys[idx%len(m.keys)]]
}
