package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main()  {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	t, _ := strconv.Atoi(strings.Split(input.Text(), "")[0])
	nums1 := make([]int, 0)
	nums2 := make([]int, 0)
	for i:=0;i<t;i++ {
		input.Scan()
		num1, _ := strconv.Atoi(strings.Split(input.Text(), " ")[0])
		num2, _ := strconv.Atoi(strings.Split(input.Text(), " ")[1])
		nums1 = append(nums1, num1)
		nums2 = append(nums2, num2)
	}
	//nums1 := []int{77, 51, 33, 69}
	//nums2 := []int{1, 2, 2, 3}
	res := adjust(nums1, nums2)
	for i:=0;i<len(res)-1;i++ {
		fmt.Printf("%d ", res[i])
	}
	fmt.Printf("%d", res[len(res)-1])
}

func adjust(nums1 []int, nums2 []int) []int {
	Map1 := make(map[int]int)
	Map2 := make([]int, 0)
	res := make([]int, 0)
	for i:= len(nums2)-1;i>=0;i-- {
		if _, ok := Map1[nums2[i]]; !ok {
			Map1[nums2[i]] = i
		} else {
			Map2 = append(Map2, nums1[i])
		}
	}
	var keys []int
	for k, _ := range Map1 {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		res = append(res, nums1[Map1[k]])
	}
	res = append(res, Map2...)
	return res
}

