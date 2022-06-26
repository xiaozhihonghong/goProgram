
/*
把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。写一个程序，输入n，打印出s的所有可能的值出现的概率。

提示：你需要用一个浮点数数组返回答案，其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 小的那个的概率。

例如
示例 1:
输入: 1
输出: [0.16667,0.16667,0.16667,0.16667,0.16667,0.16667]

示例 2:
输入: 2
输出: [0.02778,0.05556,0.08333,0.11111,0.13889,
0.16667,0.13889,0.11111,0.08333,0.05556,0.02778]
 */
package main

import (
	"fmt"
	"math"
)

func GetNums(n int) []float64 {
	res := make([]float64, 5*n+1)
	total := math.Pow(6, float64(n))
	dp := make([][]float64, n+1)
	for i:=0;i<=n;i++ {
		dp[i] = make([]float64, 6*n+1)
	}
	for i:=1;i<=6;i++ {
		dp[1][i] = 1
	}
	for i:=1;i<=n;i++ {
		for j:=i*6;j>=i;j-- {
			for k:=1;k<=6;k++ {
				if j-k >= i-1 {
					dp[i][j] += dp[i-1][j-k]
				} else {
					dp[i][j] += 0
				}
			}
			if i==n {
				res[j-n] = dp[i][j] / total
			}
		}
	}
	return res
}

func main() {
	res := GetNums(2)
	fmt.Println(res)
}
