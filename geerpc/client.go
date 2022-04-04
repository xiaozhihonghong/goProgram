package geerpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"geerpc/codec"
	"io"
	"log"
	"net"
	"sync"
)

//实现并发和异步的客户端

// Call里面包含一次RPC调用的所有消息
type Call struct {
	Seq            uint64
	ServiceMethod  string
	Args           interface{}
	Reply          interface{}
	Error          error
	Done           chan *Call   //当调用结束后，使用Done通知
}

//传入一个结束的消息给Done
func (c *Call) done()  {
	c.Done <- c
}

type Client struct {
	cc       codec.Codec
	opt      *Option
	sending  sync.Mutex   //为了保证请求有序发送
	header   codec.Header
	mu       sync.Mutex
	seq      uint64
	pending  map[uint64]*Call
	closing  bool
	shutdown bool
}

var _ io.Closer = (*Client)(nil)  //必须实现close

var ErrShutdown = errors.New("connection is shut down")

func (client *Client) Close() error {
	client.mu.Lock()
	defer client.mu.Unlock()
	if client.closing {
		return ErrShutdown
	}
	client.closing = true
	return client.cc.Close()
}

func (client *Client) IsAvailable() bool {
	client.mu.Lock()
	defer client.mu.Unlock()
	return !client.shutdown && !client.closing
}

// 注册Call到client中
func (client *Client) RegisterCall(call *Call) (uint64, error) {
	client.mu.Lock()
	defer client.mu.Unlock()
	if client.closing || client.shutdown {
		return 0, ErrShutdown
	}
	call.Seq = client.seq
	client.pending[call.Seq] = call
	client.seq++
	return call.Seq, nil
}

// 从client中移除Call,返回移除的call
func (client *Client) RemoveCall(seq uint64) *Call {
	client.mu.Lock()
	defer client.mu.Unlock()
	call := client.pending[seq]
	delete(client.pending, seq)
	return call
}

// 服务端和客户端发生错误，通知其他call,都进行关闭
func (client *Client) TerminateCall(err error) {
	client.sending.Lock()
	defer client.mu.Unlock()  //保证客户端请求的顺序
	client.mu.Lock()    // 保证每次只有一个请求
	defer client.mu.Unlock()
	client.shutdown = true
	for _, call := range client.pending {
		call.Error = err
		call.done()
	}
}

//接收响应
func (client *Client) receive()  {
	var err error
	for err == nil {
		var h codec.Header
		if err = client.cc.ReadHeader(&h); err != nil {
			break
		}
		call := client.RemoveCall(h.Seq)
		switch {
		case call == nil:
			err = client.cc.ReadBody(nil)  //call不存在
		case h.Error != "":       //call存在，服务端处理错误
			call.Error = fmt.Errorf(h.Error)
			err = client.cc.ReadBody(nil)
			call.done()
		default:
			err = client.cc.ReadBody(call.Reply)
			if err != nil {
				call.Error = errors.New("reading body " + err.Error())
			}
			call.done()   //不管是否读出，都需要关闭这次通信
		}

	}
	client.TerminateCall(err)
}

func NewClient(conn net.Conn, opt *Option) (*Client, error) {
	f := codec.NewCodecFuncMap[opt.CodeType]
	if f == nil {
		err := fmt.Errorf("invalid codec type %s", opt.CodeType)
		log.Println("rpc client:codec error:", err)
		return nil, err
	}
	//编码发送给服务端
	if err := json.NewEncoder(conn).Encode(opt); err != nil {
		log.Println("rpc client:option error: ", err)
		_ = conn.Close()
		return nil, err
	}
	return newClientCodec(f(conn), opt), nil
}

func newClientCodec(cc codec.Codec, opt *Option) *Client {
	client := &Client{
		seq: 1,
		cc: cc,
		opt: opt,
		pending: make(map[uint64]*Call),
	}
	go client.receive()
	return client
}

//创建客户端的连接
func Dial(network, address string, opts ...*Option) (client *Client, err error)  {
	opt, err := parseOptions(opts...)
	if err != nil {
		return nil, err
	}
	conn, err := net.Dial(network, address)  //创建连接
	if err != nil {
		return nil, err
	}
	defer func() {
		if client == nil {
			_ = conn.Close()
		}
	}()
	return NewClient(conn, opt)   //编码并监听接收信息
}

//保证opt正确性
func parseOptions(opts ...*Option) (*Option, error) {
	if len(opts) == 0 || opts[0] == nil {
		return DefaultOption, nil
	}
	if len(opts) != 1 {
		return nil, errors.New("number of options is more than 1")
	}
	opt := opts[0]
	opt.MagicNumber = DefaultOption.MagicNumber
	if opt.CodeType == "" {
		opt.CodeType = DefaultOption.CodeType
	}
	return opt, nil
}

//实现客服端的发送能力，可以同步也可以异步发送，异步发送的话需要使用go子协程发送，这样就实现了发送的并发，接收也是使用的go协程并发
func (client *Client) send(call *Call)  {
	client.sending.Lock()
	defer client.sending.Unlock()

	seq, err := client.RegisterCall(call)
	if err != nil {
		call.Error = err
		call.done()
		return
	}

	client.header.ServiceMethod = call.ServiceMethod
	client.header.Seq = seq
	client.header.Error = ""

	if err := client.cc.Write(&client.header, call.Args); err != nil {
		call := client.RemoveCall(seq)
		if call != nil {
			call.Error = err
			call.done()
		}
	}
}

//异步发送
func (client *Client) Go(serviceMethod string, args, reply interface{}, done chan *Call) *Call {
	if done == nil {
		done = make(chan *Call, 10)
	} else if cap(done) == 0 {
		log.Panic("rpc client: done channel is unbuffered")
	}
	call := &Call{
		ServiceMethod: serviceMethod,
		Args: args,
		Reply: reply,
		Done: done,
	}
	client.send(call)
	return call
}

//同步发送
func (client *Client) Call(serviceMethod string, args, reply interface{}) error {
	call := <- client.Go(serviceMethod, args, reply, make(chan *Call, 1)).Done  //阻塞Done,等待响应返回，同步
	return call.Error
}
