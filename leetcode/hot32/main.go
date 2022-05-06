package main

import "fmt"

// 未满足连续的情况
/*
func longestValidParenthese(s string) int {
	stack := make([]byte, 0)
	count := 0
	for i:=0;i<len(s);i++ {
		if s[i] == ')' {
			if len(stack) > 0 {
				count += 2
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, '(')
		}
	}
	return count
}
*/

//dp[i]表示以i结尾的最长有小括号的长度
func longestValidParenthese(s string) int {
	maxAns := 0
	dp := make([]int, len(s))
	for i:=1;i<len(s);i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i > 1 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i - dp[i-1] - 1 >= 0 && s[i-dp[i-1]-1] == '(' {
				if i - dp[i-1] -2 >= 0 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
			maxAns = max(maxAns, dp[i])
		}
	}
	return maxAns
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	res := longestValidParenthese(")()())")
	fmt.Println(res)
}
