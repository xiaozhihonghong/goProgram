package main

import "fmt"

//使用双指针

func maxArea(height []int) int {
	n := len(height)
	if n < 2 {
		return 0
	}
	left, right := 0, n-1
	area := 0
	for left < right {
		k := min(height[left], height[right]) * (right - left)
		area = max(k, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return area
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	height := []int{1,1}
	res := maxArea(height)
	fmt.Println(res)
}