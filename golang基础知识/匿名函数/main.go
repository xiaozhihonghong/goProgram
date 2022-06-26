package main

import (
	"fmt"
	"sort"
)
//匿名函数指的就是没有函数名称的函数，可以像一个变量一样使用，匿名函数就是作为闭包使用，可以作为回调函数使用，
//回调函数就是返回值是一个函数，在后续回调使用

func Extract(m map[string][]string) []string {     //这是正常使用的函数
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)           //这是一个匿名函数，可以先定义一个函数变量，然后实现一个函数，也可以，也可以在入参和返回值直接使用匿名函数，灵活使用
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

func main() {
	//// 匿名函数的使用:方式1
	//f1 := func(n1, n2 int) string {
	//	return strconv.Itoa(n1 + n2)
	//}
	//ret1 := f1(11, 22)
	//fmt.Println(ret1)
	//
	//// 匿名函数的使用:方式2
	//ret2, b := func(a, b int) (int, bool) {
	//	fmt.Println("哈哈")
	//	return a + b, true
	//}(11, 22)
	//fmt.Println(ret2, b)

	// 匿名函数使用：方式3-闭包
	f2 := F1()
	fmt.Println(f2(88))   //88
	fmt.Println(f2(88))   //176 s的值会保存下来，这里就是在外部调用函数内部的数据
}

func F1() func(int) int {
	// 匿名函数使用：方式3,相当于闭包：外函数的返回值是内函数的引用，内函数用到了外函数的变量
	s := 0
	return func(i int) int {
		s = s + i
		return s
	}
}

/*
闭包可以理解为一种保存函数状态的方法，当我们调用一个函数，或者执行操作，或者返回结果，总之当函数运行结束后，随即消亡，
因为函数的声明一般是在堆上，当系统检测到当前内存空间没有被引用，那么就会回收。

闭包的作用就是保存函数的运行状态，避免内存被回收。当然会占用大量的内存。

一般来说异步回调过程，会需要专门的结构来保存上下文数据，这个上下文数据一般就是异步发起的时候保存起来，
然后异步结束的时候在callback里会取出使用。支持闭包的语言，比如js，lua，golang等，你就不需要专门的结构来保存这个上下文了，
你的回调函数和函数外的变量组成了闭包，所以回调函数可以直接引用这些数据，这就是闭包的好处了

比如http里面的middlerware，闭包一个典型的优势是可以在函数范围之外，调用函数内部的变量

 */