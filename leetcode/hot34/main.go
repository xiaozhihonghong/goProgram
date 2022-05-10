package main

import "fmt"

func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 || target < nums[0] || target > nums[n-1] {
		return []int{-1, -1}
	}
	left, right := 0, n-1
	for left < right {
		mid :=left + (right - left) / 2
		if nums[mid] == target {
			left = mid
			for nums[left-1] == nums[left] {
				left--
			}
			right = mid
			for nums[right+1] == nums[right] {
				right++
			}
			return []int{left, right}
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return []int{-1, -1}
}

func main() {
	nums := []int{}
	res := searchRange(nums, 6)
	fmt.Println(res)
}
