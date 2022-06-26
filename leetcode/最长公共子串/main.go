package main

import "fmt"


func longstring(s1, s2 string) string {
	//dp[i][j]表示s1以i结尾，s2以j结尾的公共子串长度
	dp := make([][]int, 0)
	for i:=0;i<=len(s1);i++ {
		dp = append(dp, make([]int, len(s2)+1))
	}
	mmax := 0
	start := 0
	for i:=1;i<=len(s1);i++ {
		for j:=1;j<=len(s2);j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > mmax {
					mmax = dp[i][j]
					start = i - mmax
				}
			}
		}
	}
	return s1[start: start+mmax]
}

func main() {
	s1 := "aacbdef"
	s2 := "bccbdef"
	res := longstring(s1, s2)
	fmt.Println(res)
}
