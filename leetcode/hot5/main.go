package main

import "fmt"

//一般使用动态规划算法即可，时间复杂度为o(n2)

func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	//for i:=0;i<len(dp);i++ {
	//	t := make([]bool, len(s))
	//	dp[i] = t
	//}

	res := 0
	low, high := 0, 1
	for	r := 0;r<len(s);r++ {
		t := make([]bool, len(s))
		dp[r] = t
		dp[r][r] = true
		for l:=0;l<r;l++ {
			if s[l] == s[r] && (l+1>=r-1 || dp[l+1][r-1]){  //注意点一，两种情况
				dp[l][r] = true
				if r-l+1> res {
					res = r-l+1
					low = l
					high = r+1
				}
			}
		}
	}
	return s[low:high]
}

func main() {
	s := "a"
	res := longestPalindrome(s)
	fmt.Println(res)
}


