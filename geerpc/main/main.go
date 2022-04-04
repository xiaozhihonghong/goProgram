package main

import (
	"fmt"
	"geerpc"
	"log"
	"net"
	"sync"
	"time"
)

// 启动一个服务
func startServer(addr chan string)  {
	listen, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network err: ", err)
	}
	log.Println("start rpc server on", listen.Addr())
	addr <- listen.Addr().String()
	geerpc.Accept(listen)
}

func main() {
	log.SetFlags(0)
	addr := make(chan string)
	go startServer(addr)

	// 简历tcp连接,con中就是客户端发动的请求
	//conn, _ := net.Dial("tcp", <- addr)
	client, _ := geerpc.Dial("tcp", <-addr)
	defer func() {_ = client.Close()}()

	time.Sleep(time.Second)
	var wg sync.WaitGroup
	//_ = json.NewEncoder(conn).Encode(geerpc.DefaultOption)
	//cc := codec.NewGobCodec(conn)   //编解码构造函数
	for i:=0 ;i < 5;i++ {
		wg.Add(1)
		//h := &codec.Header{
		//	ServiceMethod: "Foo.Sum",
		//	Seq: uint64(i),
		//}
		//_ = cc.Write(h, fmt.Sprintf("geerpc req %d", h.Seq))  //客户端发送请求
		//_ = cc.ReadHeader(h)  //读取响应的头部
		//var reply string
		//_ = cc.ReadBody(&reply) //读取响应的body
		//log.Println("reply:", reply)
		go func(i int) {
			defer wg.Done()
			args := fmt.Sprintf("geerpc req %d", i)
			var reply string
			if err := client.Call("Foo.Sum", args, &reply); err != nil {
				log.Fatal("Call Foo.Sum error:", err)
			}
			log.Println("reply:", reply)
		}(i)
		wg.Wait()
	}
}
