package geerpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"geerpc/codec"
	"io"
	"log"
	"net"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

/*
实现服务端的通信过程
 */

const MagicNumber = 0x3bef5c

//rpc的编码头

type Option struct {
	MagicNumber int    //表示标记这个一个gorpc的请求标识
	CodeType codec.Type
	ConnectTimeout  time.Duration //连接超时
	HandleTimeout   time.Duration //客户端处发送请求、等待服务端处理请求、接受请求超时
}

//也可以选择这种初始化方式，因为都是常量

var DefaultOption = &Option{
	MagicNumber: MagicNumber,
	CodeType: codec.GobType,
	ConnectTimeout: time.Second * 10,
}

type Server struct {
	serviceMap  sync.Map
}

func NewServer() *Server {
	return &Server{}
}

var DefaultServer = NewServer()

func (server *Server) Register(rcvr interface{}) error {
	s := newService(rcvr)
	if _, dup := server.serviceMap.LoadOrStore(s.name, s); dup {
		return errors.New("rpc service already defined:" + s.name)
	}
	return nil
}

func Register(rcvr interface{}) error {
	return DefaultServer.Register(rcvr)
}

func (server *Server) findService(serviceMethod string) (svc *service, mtype *methodType, err error) {
	dot := strings.LastIndex(serviceMethod, ".")
	if dot < 0 {
		err = errors.New("rpc server:service/method request ill-formed: "+ serviceMethod)
		return
	}
	serviceName, methodName := serviceMethod[:dot], serviceMethod[dot+1:]
	svic, ok := server.serviceMap.Load(serviceName)
	if !ok {
		err = errors.New("rpc server:can't found service: "+ serviceName)
		return
	}
	svc = svic.(*service)
	mtype = svc.method[methodName]
	if mtype == nil {
		err = errors.New("rpc server:can't found method: "+ methodName)
	}
	return
}

//实现接受和处理请求,直接调用这个接口就可以启动服务
func Accept(listener net.Listener)  {
	DefaultServer.Accept(listener)
}

//实现一个accept监听请求,定义一个listen监听器，然后accept接受请求，最后handle处理请求，这里的监听器为net.Listener
func (server *Server) Accept(listener net.Listener)  {
	 for {
	 	conn, err := listener.Accept()
	 	if err != nil {
	 		log.Println("rpc server: accept error:", err)
			return
		}
		go server.Serveconn(conn)
	 }
}

//实现处理接受的请求
func (server *Server) Serveconn(conn io.ReadWriteCloser)  {
	/*
	1、使用json反序列化option实例
	2、检查里面的字段是否正确
	3、根据codetype得到编解码器
	4、serverCodec处理
	 */
	defer func() {_ = conn.Close()}()
	var opt Option
	if err := json.NewDecoder(conn).Decode(&opt); err != nil {
		log.Println("rpc server:options error:", err)
		return
	}
	f := codec.NewCodecFuncMap[opt.CodeType]
	if f == nil {
		log.Printf("rpc server: invalid codec type %s:", opt.CodeType)
		return
	}
	server.serveCodec(f(conn), &opt)
}

var invalidRequest = struct {}{}
func (server *Server) serveCodec(cc codec.Codec, opt *Option)  {
	/*
	1、接收请求
	2、处理请求
	3、发送响应
	 */
	sending := new(sync.Mutex)
	wg := new(sync.WaitGroup)
	for {
		req, err := server.readRequest(cc)
		if err != nil {
			if req == nil {
				break
			}
			req.h.Error = err.Error()
			server.sendResponse(cc, req.h, invalidRequest, sending)  //表示这是发送的一个无效的响应
		}

		wg.Add(1)
		go server.handleRequest(cc, req, sending, wg, opt.HandleTimeout)
	}
	wg.Wait()
	_ = cc.Close()
}

type request struct {
	h   *codec.Header   //head
	agrv, replyv reflect.Value   //body,请求和响应
	mtype  *methodType
	svc  *service
}

func (server *Server) readRequestHeader(cc codec.Codec) (*codec.Header, error) {
	var h codec.Header
	if err := cc.ReadHeader(&h); err != nil {
		if err != io.EOF && err != io.ErrUnexpectedEOF {
			log.Println("rpc server:read header error:", err)
		}
		return nil, err
	}
	return &h, nil
}

//实现接受请求
func (server *Server) readRequest(cc codec.Codec) (*request, error) {
	h, err := server.readRequestHeader(cc)
	if err != nil {
		return nil, err
	}
	req := &request{
		h: h,
	}
	//req.agrv = reflect.New(reflect.TypeOf(""))
	//if err = cc.ReadBody(req.agrv.Interface()); err != nil {
	//	log.Println("rpc server: read argv err:", err)
	//}
	req.svc, req.mtype, err = server.findService(h.ServiceMethod)
	if err != nil {
		return req, err
	}
	req.agrv = req.mtype.NewArgv()
	req.replyv = req.mtype.NewReplyv()
	argvi := req.agrv.Interface()
	if req.agrv.Type().Kind() != reflect.Ptr {
		argvi = req.agrv.Addr().Interface()
	}
	if err = cc.ReadBody(argvi); err != nil {
		log.Println("rpc server: read body err:", err)
		return req, err
	}
	return req, nil
}

//实现处理请求
func (server *Server) handleRequest(cc codec.Codec, req *request, mutex *sync.Mutex, wg *sync.WaitGroup, timeout time.Duration)  {
	defer wg.Done()
	//log.Println(req.h, req.agrv.Elem())
	//req.replyv = reflect.ValueOf(fmt.Sprintf("geerpc resp %d", req.h.Seq))
	called := make(chan struct{})  //代表处理没有超时，继续执行sendResponse
	sent := make(chan struct{})  //代表处理超时
	go func() {
		err := req.svc.call(req.mtype, req.agrv, req.replyv)
		called <- struct{}{}
		if err != nil {
			req.h.Error = err.Error()
			server.sendResponse(cc, req.h, invalidRequest, mutex)
			sent <- struct{}{}
			return
		}
		server.sendResponse(cc, req.h, req.replyv.Interface(), mutex)
		sent <- struct{}{}
	}()

	if timeout == 0 {   //不处理超时
		<- called
		<- sent
		return
	}
	select {
	case <- time.After(timeout):  //超时
		req.h.Error = fmt.Sprintf("rpc server:request handle timeout:expect within %s", timeout)
		server.sendResponse(cc, req.h, invalidRequest, mutex)
	case <- called:
		<- sent
	}
}

func (server *Server) sendResponse(cc codec.Codec, h *codec.Header, body interface{}, mutex *sync.Mutex)  {
	mutex.Lock()
	defer mutex.Unlock()
	if err := cc.Write(h, body); err != nil {
		log.Println("rpc server write response error:", err)
	}
}

func (s *service) call(m *methodType, argv, replyv reflect.Value) error {
	atomic.AddUint64(&m.numCalls, 1)
	f := m.method.Func
	returnValues := f.Call([]reflect.Value{s.rcvr, argv, replyv})
	if errInter := returnValues[0].Interface(); errInter != nil {
		return errInter.(error)
	}
	return nil
}


