package main

import "fmt"

//无重复最长子串
//自己开始想到的方法
func lengthOfLongestSubstring(s string) int {
	l := 0
	Map := make(map[uint8]int)
	res := 0
	r := 0
	for ; r < len(s); r++ {
		if item, ok := Map[s[r]]; ok {
			l = r
			res = max(res, r - item)
			for k, v := range Map {
				if v == item {
					delete(Map, k)
					break
				} else {
					delete(Map, k)
				}
			}
		}
		Map[s[r]] = r
	}
	if r == len(s) {
		res = max(res, r - l)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 更优的解法，map的值记录的不是下标，记录的是key出现的次数，没出现过就是0，出现过了就>0
func lengthOfLongestSubstring2(s string) int {
	var lo, hi, ret int
	r := make(map[byte]int)
	for lo <= hi && hi < len(s) {
		if cnt, ok := r[s[hi]]; ok && cnt > 0 {
			r[s[lo]]--
			lo++
			continue   //这个continue用的好，hi不会++，直接再次使用hi进行判断，看看重复的去掉没，实际和我上面的复杂度是一样的
		}
		ret = max(ret, hi-lo+1)
		r[s[hi]]++
		hi++
	}

	return ret
}


func main()  {
	s := "abcbdacbb"
	res := lengthOfLongestSubstring2(s)
	fmt.Println(res)
}