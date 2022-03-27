package main

import (
	"fmt"
	"sync"
)

//并发的map
func main() {
	//无需初始化
	var Map sync.Map

	Map.Store("chenzhi", 1)  //曾
	Map.Delete("chenzhi")  //删
	if value, ok := Map.Load("chenzhi"); ok {
		fmt.Println(value)    //查
	}

	Map.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})      //遍历
}


