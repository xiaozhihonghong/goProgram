package main

import "fmt"

// 全排列,回溯
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	list := make([]int, 0)
	res := make([][]int, 0)
	vis := make([]bool, len(nums))
	traceBack(nums, vis, list, &res)
	return res
}

func traceBack(nums []int, vis []bool, list []int, res *[][]int)  {
	for len(list) == len(nums) {
		temp := make([]int, len(list))
		copy(temp, list)
		*res = append(*res, temp)
		return
	}

	for i:=0;i<len(nums);i++ {
		if vis[i] {
			continue
		}
		vis[i] = true
		list = append(list, nums[i])
		traceBack(nums, vis, list, res)
		list = list[:len(list)-1]
		vis[i] = false
	}
}

func main() {
	nums := []int{1, 2, 3}
	res := permute(nums)
	fmt.Println(res)
}
