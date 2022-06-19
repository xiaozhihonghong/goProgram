package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	res := dfs(root)
	for i:=1;i<len(res);i++ {
		pre, cur := res[i-1], res[i]
		pre.Left, pre.Right = nil, cur
	}
}

func dfs(root *TreeNode) []*TreeNode {
	res := make([]*TreeNode, 0)
	if root==nil {
		return []*TreeNode{}
	}
	res= append(res, root)
	res = append(res, dfs(root.Left)...)
	res = append(res, dfs(root.Right)...)
	return res
}

// 空间复杂度使用o(1)
func flatten2(root *TreeNode) {
	if root == nil {
		return
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	var pre *TreeNode
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre != nil {
			pre.Left, pre.Right = nil, cur
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		pre = cur
	}
}

func main() {
	root := &TreeNode{Val:1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 6}
	flatten(root)
	for root != nil {
		fmt.Println(root.Val)
		root = root.Right
	}
}
