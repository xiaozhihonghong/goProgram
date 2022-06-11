package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := preorder[0]
	i:=0
	for ;i<len(inorder);i++ {
		if root == inorder[i] {
			break
		}
	}
	inl := inorder[:i]
	inr := inorder[i+1:]
	prel := preorder[1:len(inl)+1]
	prer := preorder[len(inl)+1:]
	root_h := &TreeNode{Val: root}
	root_h.Left = buildTree(prel, inl)
	root_h.Right = buildTree(prer, inr)
	return root_h
}

func main() {
	preorder := []int{3,9,20,15,7}
	inorder := []int{9,3,15,20,7}
	res := buildTree(preorder, inorder)
	fmt.Println(res)
}
