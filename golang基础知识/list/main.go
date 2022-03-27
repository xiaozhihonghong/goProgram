package list

import (
	"container/list"
	"fmt"
)

//list是golang实现的一个链表包，可以通过这个包实现lru

func main() {
	l := list.New()  //双向链表,初始化
	//var l2 = list.List{}  //初始化的另一种方式

	// 尾部添加
	l.PushBack("canon")
	// 头部添加
	l.PushFront(67)
	// 尾部添加后保存元素句柄
	element := l.PushBack("first")
	// 在first之后添加high
	l.InsertAfter("high", element)
	// 在first之前添加noon
	l.InsertBefore("noon", element)
	// 移除元素element
	l.Remove(element)

	//遍历双向链表
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
