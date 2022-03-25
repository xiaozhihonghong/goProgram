package gee

import (
	"net/http"
)

type handleFunc func(*Context)

//type Engine struct {
//	prefix string
//	middlerwares []handleFunc
//	router *router
//}
//
//func NewEngine() *Engine {
//	engine := &Engine{middlerwares: make([]handleFunc, 0), router: newRouter()}
//	return engine
//}
//
//func (e *Engine) Group(prefix string) *Engine {
//	newGroup := &Engine{
//		prefix: e.prefix + prefix,
//		router: e.router,
//	}
//	return newGroup
//}
//
////添加路由表,key是GET-/或GET-/hello这种形式，方法+-+path形式
//func (e *Engine) AddRouter(method string, comp string, handle handleFunc) {
//	pattern := e.prefix + comp
//	e.router.AddRouter(method, pattern, handle)
//}
//
////Get,将方法映射到路由表中，通过run方法运行
//func (e *Engine) GET(patter string, handle handleFunc)  {
//	e.AddRouter("GET", patter, handle)
//}
//
////Post
//func (e *Engine) POST(patter string, handle handleFunc)  {
//	e.AddRouter("POST", patter, handle)
//}


// gin的分布控制方式
type RouterGroup struct {
	prefix string
	middlerwares []handleFunc
	parent *RouterGroup
	engine *Engine
}

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

func NewEngine() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{
		//middlerwares: make([]handleFunc, 0),
		engine: engine,
	}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (r *RouterGroup) Group(prefix string) *RouterGroup {
	engine := r.engine
	newGroup := &RouterGroup{
		prefix: r.prefix + prefix,
		parent: r,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

//添加路由表,key是GET-/或GET-/hello这种形式，方法+-+path形式
func (r *RouterGroup) AddRouter(method string, comp string, handle handleFunc) {
	pattern := r.prefix + comp
	r.engine.router.AddRouter(method, pattern, handle)
}

//Get,将方法映射到路由表中，通过run方法运行
func (r *RouterGroup) GET(patter string, handle handleFunc)  {
	r.engine.AddRouter("GET", patter, handle)
	r.AddRouter("GET", patter, handle)        //之前一直运行结果为404，由于使用的上面的语句，单只每次从都是从engine中调用add，
														// 所以一直都是无r.prefix + comp的状态
}

//Post
func (r *RouterGroup) POST(patter string, handle handleFunc)  {
	//r.engine.AddRouter("POST", patter, handle)
	r.AddRouter("POST", patter, handle)
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

