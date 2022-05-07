package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	low, high := 0, len(nums) - 1
	for low < high {
		mid := low + (high - low) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[high] {
			if nums[mid] > target && target > nums[high]{
				high = mid - 1
			} else {
				low = mid
			}
		}
		if nums[mid] <= nums[high] {
			if nums[mid] < target && target < nums[high] {
				low = mid + 1
			} else {
				high = mid
			}
		}
	}
	return -1
}

func main() {
	nums := []int{4,5,6,7,0,1,2}
	res := search(nums, 3)
	fmt.Println(res)
}
