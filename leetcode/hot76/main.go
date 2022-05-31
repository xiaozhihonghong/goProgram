package main

import (
	"fmt"
	"math"
)

//实际上是双指针问题，不是使用滑动窗口求解
func minWindow(s string, t string) string {
	left := 0
	right := 0
	target := make(map[byte]int)
	fact := make(map[byte]int)
	match := 0
	start := 0
	end := 0
	amin := math.MaxInt32
	for i:=0;i<len(t);i++ {
		target[t[i]]++
	}
	for right < len(s) {
		// 右扩
		temp := s[right]
		right++
		if target[temp] != 0 {
			fact[temp]++
			if fact[temp] == target[temp] {
				match++
			}
		}

		// 左扩
		for match == len(target) {
			if right - left < amin {
				amin = right - left
				end = right
				start = left
			}

			c := s[left]
			left++
			if target[c] != 0 {
				//这一步需要注意，如果两者相等，表示已经算过，左边已经走完，match就不匹配了
				if fact[c] == target[c] {
					match--
				}
				fact[c]--
			}
		}
	}

	if amin == math.MaxInt32 {
		return ""
	}
	return s[start:end]
}


func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t))
}
