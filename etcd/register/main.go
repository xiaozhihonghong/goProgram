package main

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

//服务注册
type registerService struct {
	cli        *clientv3.Client   //etcd客户端
	leaseID    clientv3.LeaseID         //租约的ID
	key        string      //简单起见，使用string代替注册的服务
	value      string
	keepAlive  <- chan *clientv3.LeaseKeepAliveResponse  //租约监听
}

func NewRegisterService(key, value string, endPoints []string, ttl int64) (*registerService, error) {
	/*
		1、初始化etcd客户端
		2、服务注册，put
	 */
	client, err := clientv3.New(clientv3.Config{
		Endpoints: endPoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal("new etcd client error:", err)
	}
	service := &registerService{
		cli: client,
		key: key,
		value: value,
	}
	//注册服务
	if err = service.register(ttl);err != nil {
		return nil, err
	}
	return service, nil
}

//服务注册
func (rs *registerService) register(ttl int64) error {
	//服务注册
	log.Printf("start register and lease")
	//ctx := context.Background()
	leaseGrantResponse, err := rs.cli.Grant(context.Background(), ttl)
	if err != nil {
		return err
	}
	_, err = rs.cli.Put(context.Background(), rs.key, rs.value, clientv3.WithLease(leaseGrantResponse.ID))
	if err != nil {
		return err
	}
	// 续租定期发送心跳
	keepAliveResp, err := rs.cli.KeepAlive(context.Background(), leaseGrantResponse.ID)
	if err != nil {
		return err
	}
	rs.leaseID = leaseGrantResponse.ID
	rs.keepAlive = keepAliveResp
	log.Printf("finish register key %s, value %s and lease", rs.key, rs.value)
	return nil
}

func (rs *registerService) ListenLeaseRespChan()  {
	for leaseKeepAlive := range rs.keepAlive {
		log.Println("lease successful", leaseKeepAlive)
	}
	log.Println("lease failure, close lease")
}

func (rs *registerService) Close() error {
	//撤销租约
	if _, err := rs.cli.Revoke(context.Background(), rs.leaseID); err != nil {
		return err
	}
	log.Println("撤销租约")
	return rs.cli.Close()
}

func main() {
	endPoints := []string{"localhost:2379"}
	registerServ, err := NewRegisterService("/web/node1", "localhost:8000", endPoints, 5)
	if err != nil {
		log.Fatalln(err)
	}
	go registerServ.ListenLeaseRespChan()
	select {
	case <- time.After(20 * time.Second):
		registerServ.Close()
	}
}
