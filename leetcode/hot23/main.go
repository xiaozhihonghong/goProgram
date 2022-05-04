package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	return merge(lists)
}

func merge(lists []*ListNode) *ListNode {
	if len(lists) == 1 {
		return lists[0]
	}
	mid := len(lists) / 2
	left := merge(lists[:mid])
	right := merge(lists[mid:])
	return mergeTwoList(left, right)
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head := &ListNode{Val:-1}
	cur := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return head.Next
}

func main() {
	nums1 := []int{4,5}
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

	nums3 := []int{6}
	head3 := &ListNode{Val:2}
	q3 := head3
	for i:=0;i<len(nums3);i++ {
		t := &ListNode{Val:nums2[i]}
		q3.Next = t
		q3 = q3.Next
	}

	list := make([]*ListNode, 0)
	list = append(list, head1)
	list = append(list, head2)
	list = append(list, head3)
	newHead := mergeKLists(list)
	for newHead != nil {
		fmt.Println(newHead.Val)
		newHead = newHead.Next
	}
}
