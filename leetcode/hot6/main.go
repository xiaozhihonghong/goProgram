package main

import (
	"fmt"
)

//自己写的，模拟
func convert(s string, num int) string {
	char := make([][]byte, 0)
	for i:=0;i<len(s); {
		t := make([]byte, num)
		for j:=0;j<num;j++ {
			if i<len(s) {
				t[j] = s[i]
				i++
			} else {
				break
			}
		}
		char = append(char, t)
		for j:=num-2;j>0;j-- {
			if i <len(s) {
				t = make([]byte, num)
				t[j] = s[i]
				i++
				char = append(char, t)
			} else {
				break
			}
		}
	}

	res := ""
	for i:=0;i<num;i++ {
		for j:=0;j<len(char);j++ {
			if char[j][i] > 0 {
				res += string(char[j][i])
			}
		}
	}
	return res
}

//优化,是自己想复杂了
func convert3(s string, num int) string {
	char := make([]string, num)
	for i:=0;i<len(s); {
		for j:=0;j<num && i<len(s);j++ {
			char[j] += string(s[i])
			i++
		}

		for j:=num-2;j>0 && i<len(s);j-- {
			char[j] += string(s[i])
			i++
		}
	}

	//最后加成一列
	for j:=1;j<num;j++ {
		char[0] += char[j]
	}
	return char[0]
}


func main() {
	s := "PAYPALISHIRING"
	res := convert3(s, 4)
	fmt.Println(res)
}
