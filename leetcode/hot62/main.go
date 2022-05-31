package main

import "fmt"

func uniquePath(m, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	dp := make([][]int, m)
	for i:=0; i<m; i++ {
		dp[i] = make([]int, n)
	}
	for i:=0; i<n; i++ {
		dp[0][i] = 1
	}
	for i:=0; i<m; i++ {
		dp[i][0] = 1
	}
	for i:=1;i<m;i++ {
		for j:=1;j<n;j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func main() {
	res := uniquePath(3, 2)
	fmt.Println(res)
}
