package geeCache

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

type HttpPool struct {
	self string  //用来记录自己的地址，包括主机名/IP 和端口
	basePath string //作为节点间通讯地址的前缀，默认是 /_geecache/
	mu sync.Mutex
	peers *Map    //实现从分布式节点获取缓存值
	httpGetters map[string]*httpGetter  // keyed by e.g. "http://10.0.0.2:8008"，每个节点对应一个url
}

const (
	defaultBasePath = "/_geecache/"
	defaultReplicas = 50
)

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

//实现通信过程中获取缓存， 实现一个服务端
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
	if len(parts) != 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	name := parts[0]
	key := parts[1]
	gp, err := GetGroup(name)
	if gp == nil {
		http.Error(w, "no such group: "+ name, http.StatusNotFound)
		return
	}

	view, err := gp.GetCache(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(view.B)
}

type httpGetter struct {
	baseURL string
}

// 实现一个客户端
func (h *httpGetter) Get(group string, key string)([]byte, error)  {
	u := fmt.Sprintf(
		"%v%v/%v",
			h.baseURL,
			url.QueryEscape(group),  //QueryEscape函数对group进行转码使之可以安全的用在URL查询里
			url.QueryEscape(key),
		)                //组成一个url
	res, err := http.Get(u)  //http通过url从服务端取到了结果，核心就是这句
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()  //http包解析需要close，否则会内存泄露
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned: %v", res.Status)
	}
	bytes, err := ioutil.ReadAll(res.Body)  //读取文件或者网络请求时
	if err != nil {
		return nil, fmt.Errorf("reading response body: %v", err)
	}
	return bytes, nil
}

// 实现PeerGetter接口,
var _ PeerGetter = (*httpGetter)(nil) // _表示变量不赋值，类型为PeerGetter,
// 确保接口被实现常用的方式。即利用强制类型转换，确保 struct HTTPPool 实现了接口 PeerPicker。这样 IDE 和编译期间就可以检查，而不是等到使用的时候。


// 实例化一个一致性hash算法,peers就是一个个节点
func (p *HttpPool) Set(peers ...string)  {
	p.mu.Lock()
	defer p.mu.Unlock()
	//实例化一个hash
	p.peers = NewMap(defaultReplicas, nil)
	// 添加节点
	p.peers.Add(peers...)  //添加多个值需要加上...
	//实例化httpGetters
	p.httpGetters = make(map[string]*httpGetter, len(peers))
	//每个节点都建立和url的映射
	for _, peer := range peers {
		p.httpGetters[peer] = &httpGetter{baseURL: peer + p.basePath}
	}
}

//封装Get方法，挑选节点返回客户端
func (p *HttpPool) PickPeer(key string)(PeerGetter, bool)  {
	p.mu.Lock()
	defer p.mu.Unlock()
	//选择除自己之外的其他分布式中的节点
	if peer := p.peers.Get(key); peer != "" && peer != p.self {
		p.Log("Pick peer %s", peer)
		return p.httpGetters[peer], true
	}
	return nil, false
}

var _ PeerPicker = (*HttpPool)(nil)