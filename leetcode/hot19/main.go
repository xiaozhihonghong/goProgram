package main

import "fmt"

//删除链表的的倒数第N个节点，注意有可能删除头结点

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newHead := &ListNode{Val: -1}
	newHead.Next = head
	fast := newHead
	low := newHead
	for i:=0;i<n;i++ {
		if fast.Next == nil {
			return newHead
		}
		fast = fast.Next
	}
	for fast.Next != nil {
		low = low.Next
		fast = fast.Next
	}
	low.Next = low.Next.Next
	return newHead.Next
}

func main() {
	//nums := []int{2,3,4,5}
	head := &ListNode{Val:1}
	//p := head
	//for i:=0;i<len(nums);i++ {
	//	t := &ListNode{Val:nums[i]}
	//	p.Next = t
	//	p = p.Next
	//}
	newHead := removeNthFromEnd(head, 1)
	for newHead != nil {
		fmt.Println(newHead.Val)
		newHead = newHead.Next
	}
}
