package main

import "fmt"

// 颜色分类，使用sort，不用sort就快排
// o(N)的话使用双指针

func sortColor(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	zero := 0
	one := 0
	two := len(nums) - 1
	for one < two {
		if nums[one] == 0 {
			nums[zero], nums[one] = nums[one], nums[zero]
			zero++
			one++
		} else if nums[one] == 1 {
			one++
		} else {
			nums[two], nums[one] = nums[one], nums[two]
			two--
		}
	}
	return nums
}

func main() {
	nums := []int{0,2,1,2,0,1,0,0}
	res := sortColor(nums)
	fmt.Println(res)
}
