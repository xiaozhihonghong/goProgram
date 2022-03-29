package main

import "fmt"

// 寻找两个正序数组的中位数，要求时间复杂度为log(m+n)

//解法1，合并两个数组，找中位数，时间复杂度o(m+n)
//解法2，两个数组从前往后比较，找到中间这个数o(m+n)
//解法三，二分，每次找k/2处的数，k为第k小的数

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	if length % 2 == 1 {
		return float64(getMinK(nums1, nums2, length/2+1))   //注意点1，length/2+1
	}
	return float64(getMinK(nums1, nums2, length/2) + getMinK(nums1, nums2, length/2+1)) / 2
}

func getMinK(nums1, nums2 []int, k int) int {
	for {
		if len(nums1) == 0 {
			return nums2[k-1]
		} else if len(nums2) == 0 {
			return nums1[k-1]
		} else if k == 1 {
			return getMin(nums1[k-1], nums2[k-1])
		}
		index := getMin(k/2-1, getMin(len(nums1)-1, len(nums2)-1))
		if nums1[index] < nums2[index] {
			nums1 = nums1[index+1:]
		} else {
			nums2 = nums2[index+1:]
		}
		//k -= index + 1   //表示k = k - index - 1
		k = k - index - 1
	}
}

func getMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	res := findMedianSortedArrays(nums1, nums2)
	fmt.Println(res)
}
