package main

import "fmt"

func maximalRectangle(matrix [][]byte) int {
	mmax := 0
	height := make([]int, len(matrix[0]))
	for i:=0;i<len(matrix);i++ {
		for j:=0;j<len(matrix[0]);j++ {
			if matrix[i][j] == '1' {
				height[j] += 1
			} else {
				height[j] = 0
			}
		}
		mmax = max(mmax, largestRectangleArea(height))
	}
	return mmax
}

func largestRectangleArea(height []int) int {
	stack := make([]int, 0)
	stack = append(stack, 0)
	height = append(height, -1)
	mmax := 0
	for i:=0;i<len(height);i++ {
		for len(stack)>1&&height[i]<=height[stack[len(stack)-1]] {
			peek := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			p := stack[len(stack)-1]
			mmax = max(mmax, height[peek]*(i-p-1))
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
	matrix := [][]byte{{'1','0','1','0','0'},{'1','0','1','1','1'},{'1','1','1','1','1'},{'1','0','0','1','0'}}
	res := maximalRectangle(matrix)
	fmt.Println(res)
}
