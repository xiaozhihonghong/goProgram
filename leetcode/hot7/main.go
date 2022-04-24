package main

import (
	"fmt"
	"math"
)

func reverse(x int) (res int) {
	//if x < math.MinInt32 || x > math.MaxInt32 {
	//	return 0
	//}
	for x != 0 {
		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0
		}
		t := x % 10
		x /= 10
		res = res*10 + t
	}
	return
}

func main() {
	x := -123
	res := reverse(x)
	fmt.Println(res)
}
