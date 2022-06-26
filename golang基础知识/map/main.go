package main

import (
	"fmt"
	"sync"
)

//golang的map主要包括hmap和bmap两个结构体，hmap包括bmap的指针，即桶，桶包括原有的桶和溢出桶。扩容是指map太挤>6.5,容量编两倍
//，太松，溢出桶中删除了太多原有的元素，容量为原有的容量，重新加入新桶中。map是并发不安全的。
//并发安全的map，sync.Map。读是乐观锁，写时使用一个dirty字段写入新数据，在某个时刻方法原有字段中，做到无锁机制。读多写少的情况下使用
//否者可以直接使用加锁+map的机制控制并发，https://blog.csdn.net/a348752377/article/details/104972194
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


