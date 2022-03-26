package gee

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

func Recovery() HandleFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				c.Fail(http.StatusInternalServerError, "Internal Servel Error")
			}
		}()

		c.Next()
	}
}

func trace(message string) string {
	/*
	unsafe.Pointer只是单纯的通用指针类型，用于转换不同类型指针，它不可以参与指针运算；
	而uintptr是用于指针运算的，GC 不把 uintptr 当指针，也就是说 uintptr 无法持有对象， uintptr 类型的目标会被回收；
	unsafe.Pointer 可以和 普通指针 进行相互转换；
	unsafe.Pointer 可以和 uintptr 进行相互转换

	Caller
	获取调用函数信息
	参数作用：0表示调用函数本身

	Callers
	获取程序计算器，第二个参数会返回程序计算器列表，return值是个数

	CallersFrames
	获取栈的全部信息，通过和Callers配合来使用

	FuncForPC
	通过reflect的ValueOf().Pointer作为入参，获取函数地址、文件行、函数名等信息

	Stack
	获取栈信息
	*/
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:])

	//Builder的底层实现其实就是一个string类型的切片
	var str strings.Builder
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc) //获取对应函数
		file, line := fn.FileLine(pc)  //获取文件名和函数
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}
