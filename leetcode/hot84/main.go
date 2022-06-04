package main

import "fmt"


func largestRectangleArea(height []int) int {
	stack := make([]int, 0)
	mmax := 0
	stack = append(stack, -1)
	height = append(height,0)
	for i:=0;i<len(height);i++ {
		for len(stack) >1  && height[i] <= height[stack[len(stack)-1]] {
			//注意，我们乘的不是i这个柱子，是stack栈顶的那个柱子
			peek := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			//左边界
			l := stack[len(stack)-1]
			mmax = max(mmax, height[peek]*(i - l  - 1))
		}
		stack = append(stack, i)
	}
	return mmax
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	heights := []int{2,1,5,6,2,3}
	res := largestRectangleArea(heights)
	fmt.Println(res)
}
