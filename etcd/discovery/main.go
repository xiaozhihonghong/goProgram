package main

import (
	"context"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/clientv3"
	"log"
	"sync"
	"time"
)

type discoveryService struct {
	cli       *clientv3.Client
	services  map[string]string
	mutex     sync.Mutex
}

func NewDiscoveryService(endpoints []string) *discoveryService {
	client, err := clientv3.New(clientv3.Config{
		Endpoints: endpoints,
		DialTimeout: 5*time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	return &discoveryService{
		cli: client,
		services: make(map[string]string, 0),
	}
}

func (ds *discoveryService) discoveryFromEtcd(prefix string) error {
	resp, err := ds.cli.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		return err
	}
	for _, kvs := range resp.Kvs {
		ds.setDiscoveryToLocal(string(kvs.Key), string(kvs.Value))
	}
	//监听service的变化
	go ds.watchService(prefix)
	return nil
}

func (ds *discoveryService) setDiscoveryToLocal(key, value string) {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()
	ds.services[key] = value
}

func (ds *discoveryService) watchService(prefix string) {
	watchChan := ds.cli.Watch(context.Background(), prefix, clientv3.WithPrefix())
	log.Println("start watch")
	for wresp := range watchChan {
		for _, event := range wresp.Events {
			switch event.Type {
			case mvccpb.PUT:
				ds.setDiscoveryToLocal(string(event.Kv.Key), string(event.Kv.Value))
			case mvccpb.DELETE:
				ds.deleteService(string(event.Kv.Key))
			}
		}
	}
}

func (ds *discoveryService) deleteService(key string)  {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()
	delete(ds.services, key)
}

func (ds *discoveryService) getService() []string {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()
	services := make([]string, 0)
	for _, v := range ds.services {
		services = append(services, v)
	}
	return services
}

func (ds *discoveryService) Close() error {
	return ds.cli.Close()
}

func main() {
	/*
	1、实例化一个etcd客户端
	2、从etcd中get服务，在本地维护
	3、watch服务的变化
	4、每10秒中返回服务列表
	 */
	endPoints := []string{"127.0.0.1:2379"}
	dservices := NewDiscoveryService(endPoints)
	defer dservices.Close()
	dservices.watchService("/web/")
	dservices.watchService("/gRPC/")
	for {
		select {
		case <- time.Tick(10*time.Second):
			log.Println(dservices.getService())
		}
	}
}
