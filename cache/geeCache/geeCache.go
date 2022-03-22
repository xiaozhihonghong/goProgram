package geeCache

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

type Group struct {
	getter    Getter
	name      string
	mainCache *cache //cache是封装后的lru，Cache是没有封装的lru
	peer PeerPicker
}

var (
	mutex sync.RWMutex
	groups = make(map[string]*Group)
)

func NewGroup(name string, maxByte int64, getter Getter) *Group {
	if getter == nil {
		panic("nil Getter")
	}
	mutex.Lock()
	defer mutex.Unlock()
	g := &Group{
		getter: getter,
		name: name,
		mainCache: &cache{maxByte: maxByte},
	}
	groups[name] = g
	return g
}

//如果需要特命名空间的group
func GetGroup(name string) (*Group, error) {
	mutex.Lock()
	defer mutex.Unlock()
	if g, ok := groups[name]; ok {
		return g, nil
	}
	err := errors.New("没有这个缓存")
	return nil, err
}

//获取缓存
func (g *Group) GetCache(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("key is nil" )
	}

	if v, ok := g.mainCache.get(key); ok {
		return v, nil
	}
	// 没有缓存就从源数据中获取
	value, err := g.load(key)
	if err != nil {
		return ByteView{}, err
	}
	return value, nil
}


func (g *Group) RegisterPeers(peer PeerPicker) {
	if g.peer != nil {
		panic("RegisterPeerPicker called more than once")
	}
	g.peer = peer
}

func (g *Group) load(key string) (value ByteView, err error) {
	if g.peer != nil {
		if peer, ok := g.peer.PickPeer(key); ok {
			if value, err = g.getFromPeer(peer, key); err == nil {
				return value, nil
			}
			log.Println("[GeeCache] Failed to get from peer", err)
		}
	}

	return g.getLocally(key)
}

func (g *Group) getFromPeer(peer PeerGetter, key string) (ByteView, error) {
	bytes, err := peer.Get(g.name, key)
	if err != nil {
		return ByteView{}, err
	}
	return ByteView{bytes}, nil
}

func (g *Group) getLocally(key string) (ByteView, error) {
	v, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}
	value := ByteView{Copy(v)}
	g.mainCache.add(key, value)
	return value, nil
}
