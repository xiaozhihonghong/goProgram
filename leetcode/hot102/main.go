package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	list := make([]*TreeNode, 0)
	list = append(list, root)
	for len(list) > 0 {
		l := len(list)
		temp := make([]int, 0)
		for l > 0 {
			node := list[0]
			if node.Left != nil {
				list = append(list, node.Left)
			}
			if node.Right != nil {
				list = append(list, node.Right)
			}
			temp = append(temp, node.Val)
			list = list[1:]
			l--
		}
		res = append(res, temp)
	}
	return res
}

func main() {
	root := &TreeNode{Val:3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	res := LevelOrder(root)
	fmt.Println(res)
}
