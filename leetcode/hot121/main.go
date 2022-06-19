package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	mmax := 0
	left := 0
	for i:=1;i<len(prices);i++ {
		if prices[i] > prices[left] {
			mmax = max(mmax, prices[i] - prices[left])
		} else {
			left = i
		}
	}
	return mmax
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	prices := []int{7,1,5,3,6,4}
	res := maxProfit(prices)
	fmt.Println(res)
}
