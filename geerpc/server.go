package geerpc

import (
	"encoding/json"
	"fmt"
	"geerpc/codec"
	"io"
	"log"
	"net"
	"reflect"
	"sync"
)

/*
实现服务端的通信过程
 */

const MagicNumber = 0x3bef5c

//rpc的编码头

type Option struct {
	MagicNumber int    //表示标记这个一个gorpc的请求标识
	CodeType codec.Type
}

//也可以选择这种初始化方式，因为都是常量

var DefaultOption = &Option{
	MagicNumber: MagicNumber,
	CodeType: codec.GobType,
}

type Server struct {}

func NewServer() *Server {
	return &Server{}
}

var DefaultServer = NewServer()

//实现接受和处理请求,直接调用这个接口就可以启动服务
func Accept(listener net.Listener)  {
	DefaultServer.Accept(listener)
}

//实现一个accept监听请求,定义一个listen监听器，然后accept接受请求，最后handle处理请求，这里的监听器为net.Listener
func (s *Server) Accept(listener net.Listener)  {
	 for {
	 	conn, err := listener.Accept()
	 	if err != nil {
	 		log.Println("rpc server: accept error:", err)
			return
		}
		go s.Serveconn(conn)
	 }
}

//实现处理接受的请求
func (s *Server) Serveconn(conn io.ReadWriteCloser)  {
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
		log.Println("rpc server: invalid codec type %s:", opt.CodeType)
		return
	}
	s.serveCodec(f(conn))
}

var invalidRequest = struct {}{}
func (s *Server) serveCodec(cc codec.Codec)  {
	/*
	1、接收请求
	2、处理请求
	3、发送响应
	 */
	sending := new(sync.Mutex)
	wg := new(sync.WaitGroup)
	for {
		req, err := s.readRequest(cc)
		if err != nil {
			if req == nil {
				break
			}
			req.h.Error = err.Error()
			s.sendResponse(cc, req.h, invalidRequest, sending)  //表示这是发送的一个无效的响应
		}

		wg.Add(1)
		go s.handleRequest(cc, req, sending, wg)
	}
	wg.Wait()
	_ = cc.Close()
}

type request struct {
	h   *codec.Header   //head
	agrv, replyv reflect.Value   //body,请求和响应
}

func (s *Server) readRequestHeader(cc codec.Codec) (*codec.Header, error) {
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
func (s *Server) readRequest(cc codec.Codec) (*request, error) {
	h, err := s.readRequestHeader(cc)
	if err != nil {
		return nil, err
	}
	req := &request{
		h: h,
	}
	// todo, 目前还不知道body的类型，后续实现
	req.agrv = reflect.New(reflect.TypeOf(""))
	if err = cc.ReadBody(req.agrv.Interface()); err != nil {
		log.Println("rpc server: read argv err:", err)
	}
	return req, nil
}

//实现处理请求
func (s *Server) handleRequest(cc codec.Codec, req *request, mutex *sync.Mutex, wg *sync.WaitGroup)  {
	defer wg.Done()
	log.Println(req.h, req.agrv.Elem())
	req.replyv = reflect.ValueOf(fmt.Sprintf("geerpc resp %d", req.h.Seq))
	s.sendResponse(cc, req.h, req.replyv.Interface(), mutex)
}

func (s *Server) sendResponse(cc codec.Codec, h *codec.Header, body interface{}, mutex *sync.Mutex)  {
	mutex.Lock()
	defer mutex.Unlock()
	if err := cc.Write(h, body); err != nil {
		log.Println("rpc server write response error:", err)
	}
}


