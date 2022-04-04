package codec

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
)

//编解码接口的实现
type GobCodec struct {
	conn io.ReadWriteCloser
	buf *bufio.Writer
	dec *gob.Decoder
	enc *gob.Encoder
}

//这种写法是确保Codec接口要被实现，在IDE和编译器就能被发现，而不是使用的时候发现，本质上是强制类型转换
var _ Codec = (*GobCodec)(nil)

//需要实现Codec接口的三个方法才能成功返回Codec
func NewGobCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn)  //todo，这里con和NewWriter类型的输入不一致，为什么也可以，是由于conn包括Writer吗？
	return &GobCodec{
		conn: conn,
		buf: buf,
		dec: gob.NewDecoder(conn),
		enc: gob.NewEncoder(buf),
	}
}

func (c *GobCodec) ReadHeader(h *Header) error {
	return c.dec.Decode(h)
}

func (c *GobCodec) ReadBody(body interface{}) error {
	return c.dec.Decode(body)
}

//进行编码之后写入
func (c *GobCodec) Write(h *Header, body interface{}) error {
	defer func() {
		err := c.buf.Flush()
		if err != nil {
			_ = c.Close()
		}
	}()
	if err := c.enc.Encode(h); err != nil {
		log.Println("rpc codec:gob error encoding header:", err)
		return err
	}
	if err := c.enc.Encode(body); err != nil {
		log.Println("rpc codec:gob error encoding body:", err)
		return err
	}
	return nil
}

func (c *GobCodec) Close() error {
	return c.conn.Close()
}