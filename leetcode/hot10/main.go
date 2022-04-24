package main

import "fmt"

//正则表达式匹配

func isMatch(s, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i:=0;i<m+1;i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i:=0;i<=len(s);i++ {
		for j:=1;j<=len(p);j++ {  //p字符串不能为空串,所以从1开始
			if p[j-1] == '*' {
			 	if i > 0 && (p[j-2]==s[i-1] || p[j-2]=='.') {
			 		//匹配多个或1个或者0个
			 		dp[i][j] = dp[i-1][j] || dp[i][j-1] || dp[i][j-2]
				} else {
					//匹配0个
					dp[i][j] = dp[i][j-2]
				}
			} else {
				if i > 0 && (s[i-1] == p[j-1] || p[j-1] == '.') {
					dp[i][j] = dp[i-1][j-1]
				}
			}

		}
	}
	return dp[m][n]
}

func main() {
	res := isMatch("a", "a*")
	fmt.Println(res)
}