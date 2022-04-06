package xclient

import (
	"errors"
	"math"
	"math/rand"
	"sync"
	"time"
)

//实现的负载均衡是服务发现+客户端的实现方式，另一种方式是使用专门的负载均衡器
//这里实现服务发现，本质上就是实现了负载均衡，这里的负载均衡实际上是通过负载均衡找到通过哪个客服端实例去发送这个请求

type SelectMode int  //不同的负载均衡策略

const (
	RandomSelect  SelectMode = iota   //随机
	RoundRobinSelect                  //轮询
)

//服务发现,使用interface是因为服务发现的方式有很多种，这里提供一个统一得到接口
type Discovery interface {
	Refresh() error    //自动更新
	Update(servers []string) error  //手动更新
	Get(mode SelectMode) (string, error)
	GetAll() ([]string, error)
}

//实现一个服务发现，手工维护的
type MultiServersDiscovery struct {
	r      *rand.Rand    //随机
	mu     sync.RWMutex
	servers []string
	index   int         //轮询
}

func NewMultiServersDiscovery(servers []string) *MultiServersDiscovery {
	d := &MultiServersDiscovery{
		servers: servers,
		r: rand.New(rand.NewSource(time.Now().UnixNano())),  //使用时间戳来生成随机数
	}
	d.index = d.r.Intn(math.MaxInt32-1)  //初始化这里也是返回一个随机数
	return d
}

//必须实现服务发现接口
var _ Discovery = (*MultiServersDiscovery)(nil)

//未实现这个功能
func (d *MultiServersDiscovery) Refresh() error {
	return nil
}

func (d *MultiServersDiscovery) Update(servers []string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.servers = servers
	return nil
}

func (d *MultiServersDiscovery) Get(mode SelectMode) (string, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	n := len(d.servers)
	if n == 0 {
		return "", errors.New("rpc discovery: no available servers")
	}
	switch mode {
	case RandomSelect:
		return d.servers[d.r.Intn(n)], nil
	case RoundRobinSelect:
		s := d.servers[d.index % n]
		d.index = (d.index + 1) % n
		return s, nil
	default:
		return "", errors.New("rpc discovery: no supported select mode")
	}
}

func (d *MultiServersDiscovery) GetAll() ([]string, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	s := make([]string, len(d.servers), len(d.servers))
	copy(s, d.servers)
	return s, nil
}