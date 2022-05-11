package main

import "fmt"

// 可以通过单调栈做，时间复杂度o(n)，空间复杂度o（n）

//使用双指针解题，时间复杂度o(n),空间复杂度o（1）
func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	leftMax, rightMax := 0, 0
	left, right := 0, len(height) - 1
	res := 0
	for left < right {
		leftMax = max(height[left], leftMax)
		rightMax = max(height[right], rightMax)
		if height[left] < height[right] {
			res += leftMax - height[left]
			left++
		} else {
			res += rightMax - height[right]
			right--
		}
	}
	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	height := []int{4, 2, 0, 3, 2, 5}
	res := trap(height)
	fmt.Println(res)
}
