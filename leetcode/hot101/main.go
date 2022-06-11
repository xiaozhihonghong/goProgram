package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfs(root, root)
}

func dfs(root1 *TreeNode, root2 *TreeNode) bool {
	if root1==nil && root2==nil {
		return true
	}
	if root1==nil || root2==nil || root1.Val != root2.Val{
		return false
	}
	return dfs(root1.Left, root2.Right) && dfs(root1.Right, root2.Left)
}

func main() {
	root := &TreeNode{Val:1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 3}
	res := isSymmetric(root)
	fmt.Println(res)
}
