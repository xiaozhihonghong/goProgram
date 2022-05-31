package main

import "fmt"

// sum表示当前能走到的最大距离
func canJump (nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	sum := 0
	for i:=0;i<len(nums);i++ {
		if sum >= len(nums)-1 {
			return true
		}
		if i <= sum {
			sum = max(sum, nums[i]+i)
		}
	}
	return false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{3,2,1,0, 4}
	res := canJump(nums)
	fmt.Println(res)
}
