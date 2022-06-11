package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//方法1，中序遍历，如果非递增，则为false
//方法2，dfs，在里面判断
func isValidBST(root *TreeNode) bool {
	mmax := math.MaxInt32
	mmin := math.MinInt32
	return dfs(root, mmax, mmin)
}

// mmax和mmin代表一颗树的上下界
func dfs(root *TreeNode, mmax int, mmin int) bool {
	if root == nil {
		return true
	}
	if mmax <= root.Val || mmin >= root.Val {
		return false
	}
	return dfs(root.Left, root.Val, mmin) && dfs(root.Right, mmax, root.Val)
}

func main() {
	root := &TreeNode{Val:2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	res := isValidBST(root)
	fmt.Println(res)
}
