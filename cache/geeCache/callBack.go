package geeCache

//回调函数，从源数据获取数据

type Getter interface {
	Get(key string)([]byte, error)
}

type GetterFunc func(key string)([]byte, error)

// 这里自己不能自己改名字，需要实现Get方法
func (g GetterFunc) Get(key string) ([]byte, error) {
	return g(key)
}
