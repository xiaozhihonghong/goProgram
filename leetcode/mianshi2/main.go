package main

import "fmt"

//排序好的整数数组，从小到大，给定另一个整数X，Y, 找到在这个整数数组里面有多少个数>X, <=Y (x,Y]
//最高效的时间和空间复杂度

//小于等于X的数的个数
func GetNums(nums []int, x int) int {
	n := len(nums)
	if nums[n-1] < x {
		return n
	}
	if nums[0] > x || n == 0{
		return 0
	}
	low, high := 0, n-1
	for low < high {
		mid := (high-low)/2 + low
		if nums[mid] <= x {
			if mid+1 < n && nums[mid+1] <= x {
				low = mid + 1
			} else {
				low = mid
				break
			}
		}else {
			high = mid - 1
		}
	}
	return len(nums[:low+1])
}

//(x, y]
func GetNums2(nums []int, x, y int) int {
	n := len(nums)
	if nums[0] > y || nums[n-1] <= x || n == 0 {
		return 0
	}
	if nums[0] > x || nums[n-1] <=y {
		return n
	}
	//寻找 ，y]
	low, high := 0, n-1
	for low < high {
		mid := (high-low)/2 + low
		if nums[mid] <= y {
			if mid+1 < n && nums[mid+1] <= y {
				low = mid + 1
			} else {
				low = mid
				break
			}
		} else {
			high = mid - 1
		}
	}
	right := low
	//寻找(x，

	low, high = 0, n-1
	for low < high {
		mid := (high-low)/2 + low
		//寻找>x的最小数
		if nums[mid] > x {
			if mid -1 >=0 && nums[mid-1] <= x {
				high = mid
				break
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}
	left := high
	return len(nums[left:right+1])
}

func main() {
	nums := []int{0,1,1,1,2,3,4,5,6,7,7,8}
	res := GetNums2(nums, 1, 7)
	fmt.Println(res)
}