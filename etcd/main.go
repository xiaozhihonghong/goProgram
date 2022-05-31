package main

import "fmt"

func minDistance(word1, word2 string) int {
	m := len(word1)
	n := len(word2)
	// 第0为表示空字符串
	dp := make([][]int, m+1)
	for i:=0;i<=m;i++ {
		dp[i] = make([]int, n+1)
	}
	for i:=0;i<=m;i++ {
		dp[i][0] = i
	}
	for j:=0;j<=n;j++ {
		dp[0][j] = j
	}
	for i:=1;i<=m;i++ {
		for j:=1;j<=n;j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			}
		}
	}
	return dp[m][n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	word1 := "horse"
	word2 := "ros"
	res := minDistance(word1, word2)
	fmt.Println(res)
}
