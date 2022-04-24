package main

import (
	"fmt"
)


func main()  {
	//input := bufio.NewScanner(os.Stdin)
	//input.Scan()
	//n, _ := strconv.Atoi(strings.Split(input.Text(), " ")[0])
	//m, _ := strconv.Atoi(strings.Split(input.Text(), " ")[1])
	//if n > m {
	//	n, m = m, n
	//}
	res := adjust(5, 8, 0)
	fmt.Println(res)
}

var count1, count2, count3 int

func adjust(a, b, count int) int {
	if a!=b {
		if a > b {
			a, b = b, a
		}
		if a * 2 <= b {
			count1 = adjust(a*2, b, count)
		}
		count2 = adjust(a, b/2, count)
		count3 = adjust(a+1, b, count)
		count = min(count1, min(count2, count3)) + 1
	}
	return count
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
