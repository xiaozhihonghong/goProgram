package main

import "fmt"

//dp的方法
func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i:=1;i<=len(s);i++ {
		for j:=0;j<i;j++ {
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

//超时
func dfs(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	for i:=0;i<len(wordDict);i++ {
		if len(s) >= len(wordDict[i]) && s[0:len(wordDict[i])] == wordDict[i] {
			return dfs(s[len(wordDict[i]):], wordDict)
		}
	}
	return false
}

func main() {
	s := "applepenapple"
	wordDict := []string{"apple", "pen"}
	res := wordBreak(s, wordDict)
	fmt.Println(res)
}
