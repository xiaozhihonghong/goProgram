package gee

import (
	"net/http"
	"path"
	"strings"
	"text/template"
)

type HandleFunc func(*Context)

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
	middlerwares []HandleFunc
	parent *RouterGroup
	engine *Engine
}

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup

	htmlTempaltes *template.Template   //将模板加载进内存
	funcMap template.FuncMap   //自定义函数渲染
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
func (r *RouterGroup) AddRouter(method string, comp string, handle HandleFunc) {
	pattern := r.prefix + comp
	r.engine.router.AddRouter(method, pattern, handle)
}

//Get,将方法映射到路由表中，通过run方法运行
func (r *RouterGroup) GET(patter string, handle HandleFunc)  {
	//r.engine.AddRouter("GET", patter, handle)
	r.AddRouter("GET", patter, handle)        //之前一直运行结果为404，由于使用的上面的语句，单只每次从都是从engine中调用add，
														// 所以一直都是无r.prefix + comp的状态
}

//Post
func (r *RouterGroup) POST(patter string, handle HandleFunc)  {
	//r.engine.AddRouter("POST", patter, handle)
	r.AddRouter("POST", patter, handle)
}

func (r *RouterGroup) Use(middleWares ...HandleFunc)  {
	r.middlerwares = append(r.middlerwares, middleWares...)
}

//run
func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

// run里面的ListenAndServe()里面的handler是一个接口，必须实现里面的HttpServe方法
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	//c := NewContext(w, r)
	//e.router.handle(c)

	var middleWares []HandleFunc  //两个地方添加中间件，一个是该group中所有的公共中间件，一个是对应路由上特定的中间件
	for _, group := range e.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			//看url是哪个group中的，找到该group中的所有中间件
			middleWares = append(middleWares, group.middlerwares...)
		}
	}
	c := NewContext(w, r)
	c.handlers = middleWares
	c.engine = e
	e.router.handle(c)
}

// http.FileSystem只有一个open接口，需要实现
func (r *RouterGroup) CreateStaticHandler(relativePath string, system http.FileSystem) HandleFunc {
	abPath := path.Join(r.prefix, relativePath)  //拼接路径
	//http.FileServer 返回的Handler将会进行查找，并将与文件夹或文件系统有关的内容以参数的形式返回给你
	fileServer := http.StripPrefix(abPath, http.FileServer(system))  //将请求定向到你通过参数指定的请求处理处之前，将特定的prefix从URL中过滤出去。
	return func(c *Context) {
		file := c.Param("filepath")
		if _, err := system.Open(file);err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(c.W, c.Req)  //将fileServer路径文件解析
	}
}

func (r *RouterGroup) Static(relativePath string, root string)  {
	handler := r.CreateStaticHandler(relativePath, http.Dir(root))  //渲染
	urlPattern := path.Join(relativePath, "/*filepath")
	r.GET(urlPattern, handler)
}

func (e *Engine) SetFuncMap(funcMap template.FuncMap)  {
	e.funcMap = funcMap
}

func (e *Engine) LoadHTMLGlob(pattern string)  {
	e.htmlTempaltes = template.Must(template.New("").Funcs(e.funcMap).ParseGlob(pattern))
}

