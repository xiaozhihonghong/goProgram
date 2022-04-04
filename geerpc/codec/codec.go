package codec

import "io"

//封装header
type Header struct {
	ServiceMethod string  //服务和方法
	Seq uint64  //客户端的序号
	Error string   //自定义错误信息
}

//定义编解码接口，接口是为了实现不同的编码实例，比如json和gob，
//gob是Golang包自带的一个数据结构序列化的编码/解码工具。编码使用Encoder，解码使用Decoder。一种典型的应用场景就是RPC
type Codec interface {
	io.Closer    //一个接口，不需要指针，关闭io, 这里的Close()接口也必须实现
	ReadHeader(*Header) error  //header和body是解码
	ReadBody(interface{}) error  //body类型可能不一样，使用interface
	Write(*Header, interface{}) error   //这是编码，相当于编码写入之后输出响应
}

//todo，定义一个新的类型，本质上Type就是string，所以这样定义有什么好处？，本质上和把Type定义为struct是同一种方式
//可能事项吧下面的map的key定义死为json和gob两种形式，防止输入其他的错误
type Type string

const (
	GobType Type = "application/gob"
	JsonType Type = "application/json"
)

type NewCodecFunc func(io.ReadWriteCloser) Codec

var NewCodecFuncMap map[Type]NewCodecFunc

//返回一个构造函数，可以通过map的Type得到相应的构造函数，也就是初始化函数，进而构建实例
func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
