package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	// 原来有的
	Req *http.Request
	W http.ResponseWriter
	//request包含的路径和方法
	Method string
	Path   string
	//响应中包含的状态码
	StatusCode int
}

func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Req:    req,
		W:      w,
		Method: req.Method,
		Path:   req.URL.Path,
	}
}

// 根据key从post，put表单中查询参数, todo,暂时还不知道有什么用
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

//查询req的url内容， todo，暂时还不知道有什么用
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// 响应报文起始行需要状态码，构造响应的状态码
func (c *Context) Status(code int)  {
	c.StatusCode = code
	c.W.WriteHeader(code)
}

//响应报文首部是键值对，写入响应报文的键值对
func (c *Context) SetHeader(key string, value string)  {
	c.W.Header().Set(key, value)
}

// 构造字符串形式的响应
func (c *Context) String(code int, format string, values ...interface{})  {
	c.Status(code)
	c.SetHeader("Content-type", "text/plain")  //内容类型
	c.W.Write([]byte(fmt.Sprintf(format, values...))) //body的信息
}

//构造json格式信息，包括头部和body
func (c *Context) JSON(code int, obj interface{})  {
	c.Status(code)
	c.SetHeader("Content-type", "application/json")
	encoder := json.NewEncoder(c.W) //返回一个新的编码器写入w
	if err := encoder.Encode(obj); err != nil {
		//http.Error(c.w, err.Error(), 500)
		panic(err)
		return
	}
}

// 直接以[]byte形式返回响应数据
func (c *Context) Data(code int, data []byte)  {
	c.Status(code)
	c.W.Write(data)
}

//构造HTML形式数据
func (c *Context) HTML(code int, html string)  {
	c.Status(code)
	c.SetHeader("Content-type", "text/html")
	c.W.Write([]byte(html))
}

//todo，后面还有其他形式的数据，有需要还可以补充


