package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	for i:=0;i<len(nums)-2;i++ {
		if nums[i] > 0 {
			continue
		}
		if i>0 && (nums[i]==nums[i-1]) {
			continue
		}
		low, high := i+1, len(nums)-1
		for low < high {
			if nums[i] + nums[low] + nums[high] == 0 {
				temp := []int{nums[i], nums[low], nums[high]}
				res = append(res, temp)
				for low < high && nums[low]== nums[low+1] {
					low++
				}
				for low < high && nums[high]== nums[high-1] {
					high--
				}
				low++
				high--
			} else if nums[i] + nums[low] + nums[high] < 0 {
				low++
			} else {
				high--
			}
		}
	}
	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	fmt.Println(res)
}
