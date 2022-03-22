package geeCache

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type HttpPool struct {
	self string  //用来记录自己的地址，包括主机名/IP 和端口
	basePath string //作为节点间通讯地址的前缀，默认是 /_geecache/
}

const defaultBasePath = "/_geecache/"

// 以后初始化命名在struct前面加个New
func NewHttpPool(self string) *HttpPool {
	return &HttpPool{
		self:     self,
		basePath: defaultBasePath,
	}
}

//log函数
func (p *HttpPool) Log(format string, v ...interface{})  {
	log.Printf("[Server %s] %s", p.self, fmt.Sprintf(format, v...))
}

//实现通信过程中获取缓存
//当我们需要修改结构体的变量内容的时候，方法传入的结构体变量参数需要使用指针，也就是结构体的地址。
//如果仅仅是读取结构体变量，可以不使用指针，直接传递引用即可。
//http.ResponseWriter是一个接口，不需要指针
//这里是实现main函数中http.ListenAndServe的handler接口，不能随便改名字
func (p *HttpPool) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	/**、
	1.判断前缀是不是默认的
	2.判断路径是不是groupname和key两个部分
	3.从group中获取缓存
	4.将缓存写入body返回用户
	 */
	if !strings.HasPrefix(r.URL.Path, p.basePath) {
		panic("没有路径前缀")
	}
	p.Log("%s %s", r.Method, r.URL.Path)
	parts := strings.SplitN(r.URL.Path[len(p.basePath):], "/", 2)
	name := parts[0]
	key := parts[1]
	gp, err := GetGroup(name)
	if err != nil {
		http.Error(w, err.Error() + name, http.StatusNotFound)
		return
	}

	view, err := gp.GetCache(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(view.b)
}
