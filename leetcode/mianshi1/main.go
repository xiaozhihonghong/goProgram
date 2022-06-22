package main

import "fmt"

// 给定开始日期，结束日期，例如2022-03-15，2022-06-06，可以认为参数已经拆分好年月日。
// 按照自然月切分，每个自然月输出一行，包含该自然月的开始日期和结束日期。

func main() {
	year1, year2 := 2022, 2024
	month1, month2 := 3, 6
	day1, day2 := 15, 6
	if year1>year2 {
		fmt.Println("false year!")
	}
	hashMap := make(map[int]int, 0)
	for i:=0;i<12;i++ {
		if i==1 || i==3 || i==5 || i==7 || i==8 || i==10 || i==12 {
			hashMap[i] = 31
		} else if i==2 {
			hashMap[i] = 28
		} else {
			hashMap[i] = 30
		}
	}

	for i:=year1;i<=year2;i++ {
		//闰年
		if i%400 == 0 || (i%4 == 0 && i%100 != 0) {
			hashMap[2] = 29
		}
		if year1 == year2 {
			//年份相同
			for j := month1; j <= month2; j++ {
				//第一天和最后一天需要单独判断
				if j == month1 {
					fmt.Printf("year=%d month=%d day=%d, ", i, j, day1)
					fmt.Printf("year=%d month=%d day=%d\n", i, j, hashMap[j])
				} else if j == month2 {
					fmt.Printf("year=%d month=%d day=%d, ", i, j, 1)
					fmt.Printf("year=%d month=%d day=%d\n", i, j, day2)
				} else {
					fmt.Printf("year=%d month=%d day=%d, ", i, j, 1)
					fmt.Printf("year=%d month=%d day=%d\n", i, j, hashMap[j])
				}
			}
		} else {
			//年份不同
			if i == year1 {
				//首年
				for j := month1; j <= 12; j++ {
					if j == month1 {
						fmt.Printf("year=%d month=%d day=%d, ", i, j, day1)
						fmt.Printf("year=%d month=%d day=%d\n", i, j, hashMap[j])
					} else {
						fmt.Printf("year=%d month=%d day=%d, ", i, j, 1)
						fmt.Printf("year=%d month=%d day=%d\n", i, j, hashMap[j])
					}
				}
			} else if i == year2 {
				//最后一年
				for j := 1; j <= month2; j++ {
					if j == month2 {
						fmt.Printf("year=%d month=%d day=%d, ", i, j, 1)
						fmt.Printf("year=%d month=%d day=%d\n", i, j, day2)
					} else {
						fmt.Printf("year=%d month=%d day=%d, ", i, j, 1)
						fmt.Printf("year=%d month=%d day=%d\n", i, j, hashMap[j])
					}
				}
			} else {
				// 中间年份
				for j := 1; j <= 12; j++ {
					fmt.Printf("year=%d month=%d day=%d, ", i, j, 1)
					fmt.Printf("year=%d month=%d day=%d\n", i, j, hashMap[j])
				}
			}
		}
	}
}
