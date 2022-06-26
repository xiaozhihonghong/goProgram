package main

import "fmt"

func singleNumber(nums []int) int {
	res := nums[0]
	for i:=1;i<len(nums);i++ {
		res = res ^ nums[i]
	}
	return res
}

func main() {
	nums := []int{4,1,2,1,2}
	res := singleNumber(nums)
	fmt.Println(res)
}
