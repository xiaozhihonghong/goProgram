package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	head := &ListNode{Val:-1}
	p := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}
		p = p.Next
	}
	if list1 != nil {
		p.Next = list1
	}
	if list2 != nil {
		p.Next = list2
	}
	return head.Next
}

func main() {
	nums1 := []int{2,4}
	head1 := &ListNode{Val:1}
	p := head1
	for i:=0;i<len(nums1);i++ {
		t := &ListNode{Val:nums1[i]}
		p.Next = t
		p = p.Next
	}

	nums2 := []int{3,4}
	head2 := &ListNode{Val:1}
	q := head2
	for i:=0;i<len(nums2);i++ {
		t := &ListNode{Val:nums2[i]}
		q.Next = t
		q = q.Next
	}
	newHead := mergeTwoLists(head1, head2)
	for newHead != nil {
		fmt.Println(newHead.Val)
		newHead = newHead.Next
	}
}