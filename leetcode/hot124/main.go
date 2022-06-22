package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
type Result struct {
	Single int  //左子树或者右子树贡献的路径值
	Path   int  //路径的最大值
}

func maxPathSum(root *TreeNode) int {
	if root==nil {
		return 0
	}
	return dfs(root).Path
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func dfs(root *TreeNode) *Result {
	if root == nil {
		return &Result{
			Single: 0,
			Path: 0,
		}
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	res := &Result{}
	if left.Single > right.Single {
		res.Single = max(left.Single+root.Val, 0)
	} else {
		res.Single = max(right.Single+root.Val, 0)
	}
	maxPath := max(left.Path, right.Path)
	res.Path = max(maxPath, left.Single+root.Val+right.Single)
	return res
}

func main() {
	root := &TreeNode{Val:-10}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	res := maxPathSum(root)
	fmt.Println(res)
}
