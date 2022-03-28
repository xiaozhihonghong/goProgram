package main

//星环科技笔试题1， 一个数据，两两数据交换，使用最少次数使其有序
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	t, _ := strconv.Atoi(strings.Split(input.Text(), "")[0])
	for i:=0;i<t;i++ {
		input.Scan()
		n, _ := strconv.Atoi(strings.Split(input.Text(), "")[0])
		input.Scan()
		nums := make([]int, 0)
		for j:=0;j<n;j++ {
			num, _ := strconv.Atoi(strings.Split(input.Text(), " ")[j])
			nums = append(nums, num)
		}
		res := adjust(nums)
		fmt.Println(res)
	}
	//nums := []int{2, 3, 5, 4, 1}
	//res := adjust(nums)
	//fmt.Println(res)
}

func adjust(nums []int) int {
	res := 0
	for i:=1;i<len(nums);i++ {
		if i != nums[i-1] {

			for j:=i+1;j<=len(nums);j++ {
				if nums[j-1]==i {
					nums[i-1], nums[j-1] = nums[j-1], nums[i-1]
					res++
					break
				}
			}
		}
	}
	return res
}
