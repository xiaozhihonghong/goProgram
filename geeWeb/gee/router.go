package gee

import (
	"net/http"
	"strings"
)

type router struct {
	roots map[string]*node
	handler map[string]handleFunc
}

func newRouter() *router {
	return &router{
		roots: make(map[string]*node),
		handler: make(map[string]handleFunc),
	}
}

//将一个pattern按/划分
func parsePattern(pattern string) []string {
	p := strings.Split(pattern, "/")

	parts := make([]string, 0)

	for _, item := range p {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

//添加路由表,key是GET-/或GET-/hello这种形式，方法+-+path形式, root是以method为key，pattern为node的map
func (r *router) AddRouter(method string, pattern string, handle handleFunc) {
	parts := parsePattern(pattern)
	//log.Printf("route %4s - %s", method, pattern)
	key := method + "-" + pattern
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handler[key] = handle
}

// 查找路由,返回node和动态替换的值，key为通配符，value为真实替换的值
func (r *router) getRouter(method string, path string) (*node, map[string]string)  {
	searchParts := parsePattern(path)
	params := make(map[string]string)

	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		// 返回替换的值，也就是params
		parts := parsePattern(n.pattern)
		for idx, part := range parts {
			if part[0] == ':' {
				// 通配符，:替换当前的一个值
				params[part[1:]] = searchParts[idx]
			}
			if part[0] == '*' && len(part) > 1 {
				// *替换后面所有的值
				params[part[1:]] = strings.Join(searchParts[idx:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *router) getRouters(method string) []*node {
	root , ok := r.roots[method]
	if !ok {
		return nil
	}
	nodes := make([]*node, 0)
	root.travel(&nodes)
	return nodes
}

//解析请求的路径
func (r *router) handle(c *Context)  {
	n, params := r.getRouter(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern   //key变成了替换后的值
		r.handler[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
	//key := c.Method + "-" + c.Path
	//if handler, ok := r.handler[key];ok {
	//	handler(c)   //注册的处理方法, todo 这里还有点疑惑，注册方法应该是还没实现
	//} else {
	//	c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path) //todo，实现String只是为了报错吗？
	//}
}
