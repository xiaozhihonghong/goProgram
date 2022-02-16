package goProgram

import "fmt"

func MyEcho(yzj interface{}) {
    fmt.Printf("yzj is of type %T\n", yzj)
 }

 func main() {
     var yzj interface{}    // 空接口的使用，空接口类型的变量可以保存任何类型的值,空格口类型的变量非常类似于弱类型语言中的变量，未被初始化的interface默认初始值为nil。
     MyEcho(yzj)
     yzj = 100                //将其类型定义为“int”
     MyEcho(yzj)
     yzj = "Golang"            //将其类型定义为“string”
    MyEcho(yzj)
 }