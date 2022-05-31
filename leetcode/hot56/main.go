package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	//k := 0
	//l := 0
	res := make([][]int, 0)
	for i:=1;i<len(intervals);i++ {
		//if intervals[i][0] <= intervals[k][1] {
		//	k = i
		//} else {
		//	res = append(res, []int{intervals[l][0], intervals[i-1][1]})
		//	if i < len(intervals) -1 {
		//		l = i + 1
		//	}
		//}
		left := intervals[i][0]
		right := intervals[i][1]
		for i<len(intervals)-1 && intervals[i+1][0]<=right {
			i++
			right = max(right, intervals[i][1])
		}
		res = append(res, []int{left, right})
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	intervals := [][]int{{1,3}, {2,6}, {8,10}, {15,18}}
	res := merge(intervals)
	fmt.Println(res)
}
