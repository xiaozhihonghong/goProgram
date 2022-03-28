package main

import "fmt"

//两数相加

type ListNode struct {
	Val int
	Next *ListNode
}

// 第一次想到的思路，反转链表，然后直接相加
// 理解错误，返回一个逆序的链表
func Reverse(node *ListNode) *ListNode {
	var pre *ListNode
	cur := node
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

//看错题目，返回的是一个链表，这里表示直接返回了他们的和
func addTwoNumbers(rl1 *ListNode, rl2 *ListNode) int {
	//rl1 := Reverse(l1)
	//rl2 := Reverse(l2)
	res := 0
	t := 0
	idx := 1
	for rl1 != nil && rl2 != nil {
		mp := rl1.Val+ rl2.Val + t
		if mp > 9 {
			res = res + (mp - 10) * idx
			t = 1
		} else {
			res = res + mp * idx
			t = 0
		}
		idx = idx * 10
		rl1 = rl1.Next
		rl2 = rl2.Next
		if rl1==nil && rl2==nil && t == 1 {
			res = res + t * idx
		}
	}
	for rl1 != nil {
		res += rl1.Val * idx
		idx = idx * 10
		rl1 = rl1.Next
	}
	for rl2 != nil {
		res += rl2.Val * idx
		idx = idx * 10
		rl2 = rl2.Next
	}
	return res
}

//返回逆序的链表
func addTwoNumbers2(rl1 *ListNode, rl2 *ListNode) *ListNode {
	t := 0 //进位
	var head *ListNode
	var cur *ListNode
	for rl1 != nil && rl2 != nil {
		res := (rl1.Val + rl2.Val + t) % 10
		t = (rl1.Val + rl2.Val + t) / 10
		if head == nil {
			head = &ListNode{Val: res}
			cur = head
		} else {
			cur.Next = &ListNode{Val: res}
			cur = cur.Next
		}
		rl1 = rl1.Next
		rl2 = rl2.Next
		if rl1==nil && rl2==nil && t == 1 {
			cur.Next = &ListNode{Val: t}
			cur = cur.Next
		}
	}
	for rl1 != nil {
		cur.Next = &ListNode{Val: rl1.Val}
		cur = cur.Next
		rl1 = rl1.Next
	}
	for rl2 != nil {
		cur.Next = &ListNode{Val: rl2.Val}
		cur = cur.Next
		rl2 = rl2.Next
	}
	return head
}

func main() {
	nums1 := []int{4, 7}
	l1 := &ListNode{Val: 2}
	cur1 := l1
	for _, num := range nums1 {
		cur1.Next = &ListNode{Val: num}
		cur1 = cur1.Next
	}

	nums2 := []int{6, 4}
	l2 := &ListNode{Val: 5}
	cur2 := l2
	for _, num := range nums2 {
		cur2.Next = &ListNode{Val: num}
		cur2 = cur2.Next
	}
	res := addTwoNumbers2(l1, l2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
	//fmt.Println(res)
}
