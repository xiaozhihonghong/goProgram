package main

import "fmt"

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	for j:=len(nums)-1;j>i;j-- {
		if nums[i] < nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	reverse(nums[i+1:])
}

func reverse(nums []int)  {
	n := len(nums)
	for i:=0;i<n/2;i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
}

func main() {
	nums := []int{1,1,5}
	nextPermutation(nums)
	fmt.Println(nums)
}
