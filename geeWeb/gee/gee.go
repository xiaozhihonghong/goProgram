package gee

import (
	"net/http"
)

type handleFunc func(*Context)

type Engine struct {
	router *router
}

func NewEngine() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

//添加路由表,key是GET-/或GET-/hello这种形式，方法+-+path形式
func (e *Engine) AddRouter(method string, pattern string, handle handleFunc) {
	e.router.AddRouter(method, pattern, handle)
}

//Get,将方法映射到路由表中，通过run方法运行
func (e *Engine) GET(patter string, handle handleFunc)  {
	e.AddRouter("GET", patter, handle)
}

//Post
func (e *Engine) POST(patter string, handle handleFunc)  {
	e.AddRouter("POST", patter, handle)
}

//run
func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

// run里面的ListenAndServe()里面的handler是一个接口，必须实现里面的HttpServe方法
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	c := NewContext(w, r)
	e.router.handle(c)
}

