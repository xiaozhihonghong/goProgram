package main

import "fmt"

//使用dp

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = max(nums[0], 0)
	m := dp[0]
	for i:=1;i<len(nums);i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		m = max(m, dp[i])
	}
	return m
}

func max(x, y int) int {
	for x > y {
		return x
	}
	return y
}
func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res := maxSubArray(nums)
	fmt.Println(res)
}
