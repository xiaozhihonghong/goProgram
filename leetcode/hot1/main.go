package main

import "fmt"

func twoSum(nums []int, target int) []int {
	exsit := make(map[int]int)
	for i:=0;i<len(nums);i++ {
		if j, ok := exsit[target-nums[i]]; ok {
			return []int{i, j}
		} else {
			exsit[nums[i]] = i
		}
	}
	return []int{}
}

func main()  {
	nums := []int{3, 2, 4}
	target := 6
	res := twoSum(nums, target)
	fmt.Println(res)
}
