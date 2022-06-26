package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	hashMap := make(map[int]bool)
	for _, num := range nums {
		hashMap[num] = true
	}

	longStack := 0
	for num := range hashMap {
		if !hashMap[num-1] {
			curmax := 1
			curNum := num
			for hashMap[curNum+1] {
				curmax++
				curNum++
			}
			if longStack < curmax {
				longStack = curmax
			}
		}
	}
	return longStack
}

func main() {
	nums := []int{100,4,200,1,3,2}
	res := longestConsecutive(nums)
	fmt.Println(res)
}
