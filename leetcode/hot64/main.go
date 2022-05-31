package main

import "fmt"

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	if m <= 0 || n <=0 {
		return 0
	}
	dp := make([][]int, m)
	for i:=0;i<m;i++ {
		dp[i] = make([]int, n)
		for j:=0;j<n;j++ {
			if i == 0 && j==0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j==0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}
	return dp[m-1][n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	grid := [][]int{{1,2,3},{4,5,6}}
	res := minPathSum(grid)
	fmt.Println(res)
}
