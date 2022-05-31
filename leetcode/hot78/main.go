package main

import "fmt"

func subSets(nums []int) [][]int {
	res := make([][]int, 0)
	list := make([]int, 0)
	dfs(nums, 0, list, &res)
	return res
}

func dfs(nums []int, index int, list []int, res *[][]int) {
	temp := make([]int, len(list))
	copy(temp, list)
	*res = append(*res, temp)
	for i:=index;i<len(nums);i++ {
		list = append(list, nums[i])
		dfs(nums, i+1, list, res)
		list = list[:len(list)-1]
	}
}

func main() {
	nums := []int{1,2,3}
	res := subSets(nums)
	fmt.Println(res)
}
