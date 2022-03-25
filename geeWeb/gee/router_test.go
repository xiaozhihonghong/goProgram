package gee

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.AddRouter("GET", "/", nil)
	r.AddRouter("GET", "/hello/:name", nil)
	r.AddRouter("GET", "/hello/b/c", nil)
	r.AddRouter("GET", "/hi/:name", nil)
	r.AddRouter("GET", "/assert/*filePath", nil)
	return r
}

//Test之后的首字母必须大写
func TestParsePattern(t *testing.T)  {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test parsePattern failed!")
	}
}

// todo 测试还有问题
func TestGetRouter(t *testing.T)  {
	r := newTestRouter()
	n, ps := r.getRouter("GET", "/hello/geektutu")

	if n == nil {
		t.Fatal("router == nil")
	}

	if n.pattern != "/hello/:name" {
		t.Fatal("should be /hello/:name")
	}

	if ps["name"] != "geektutu" {
		t.Fatal("name should be geektutu")
	}
}

func TestGetRoutes(t *testing.T) {
	r := newTestRouter()
	nodes := r.getRouters("GET")
	for i, n := range nodes {
		fmt.Println(i+1, n)
	}
	if len(nodes) != 5 {
		t.Fatal("should be 4")
	}
}
