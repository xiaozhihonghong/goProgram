package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 声明整型变量a并赋初值
	var a int = 1024
	// 获取变量a的反射值对象(a的地址)
	valueOfA := reflect.ValueOf(&a)
	// 取出a地址的元素(a的值)
	valueOfA = valueOfA.Elem()
	// 修改a的值为1
	valueOfA.SetInt(1)
	// 打印a的值
	fmt.Println(valueOfA.Int())
}
