package gee

import (
	"log"
	"net/http"
)
type handleFunc func(*Context)

type router struct {
	handler map[string]handleFunc
}

func newrouter() *router {
	return &router{
		handler: make(map[string]handleFunc),
	}
}

//添加路由表,key是GET-/或GET-/hello这种形式，方法+-+path形式
func (r *router) AddRouter(method string, pattern string, handle handleFunc) {
	log.Printf("route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handler[key] = handle
}

//解析请求的路径
func (r *router) handle(c *Context)  {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handler[key];ok {
		handler(c)   //注册的处理方法, todo 这里还有点疑惑，注册方法应该是还没实现
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path) //todo，实现String只是为了报错吗？
	}
}
