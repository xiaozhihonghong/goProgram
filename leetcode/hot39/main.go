package main

import (
	"fmt"
)

func combinationSum(nums []int, target int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res := make([][]int, 0)
	temp := make([]int, 0)
	traceBack(nums, 0, target, 0, temp, &res)
	return res
}

//var res [][]int

func traceBack(nums []int, sum, target, index int, temp []int, res *[][]int) {
	if sum == target {
		t := make([]int, len(temp))
		copy(t, temp)
		*res = append(*res, t)
	}
	for i:=index;i<len(nums);i++ {
		if target - sum >= nums[i] {
			sum += nums[i]
			temp = append(temp, nums[i])
			traceBack(nums, sum, target, i, temp, res)
			temp = temp[:len(temp)-1]
			sum -= nums[i]
		}
	}
}

func main() {
	nums := []int{2, 6, 3, 7}
	r := combinationSum(nums, 7)
	fmt.Println(r)
}
