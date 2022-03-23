package singleFlight

import "sync"

type call struct {
	wg sync.WaitGroup
	val interface{}
	err error
}

type Group struct {
	mutex sync.Mutex
	m map[string]*call
}

func (g *Group) Do(key string, fn func()(interface{}, error)) (interface{}, error) {
	g.mutex.Lock()
	//初始化
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	//如果已经有完成的请求，直接返回
	if v, ok := g.m[key];ok {
		g.mutex.Unlock()
		v.wg.Wait()  //有请求进行中，等待
		return v.val, nil
	}

	//添加一个进行中的请求
	v := new(call)
	v.wg.Add(1)
	g.m[key] = v
	g.mutex.Unlock()

	//发起请求
	v.val, v.err = fn()
	v.wg.Done()

	//更新请求
	g.mutex.Lock()
	delete(g.m, key)
	g.mutex.Unlock()

	return v.val, v.err
}
