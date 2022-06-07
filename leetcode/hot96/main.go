package main

import "fmt"

func dfs(num int, ) int {
	hashMap := make(map[int]int, 0)
	return numTrees(num, hashMap)
}

//最容易想到的解法，递归求解
//为了降低时间复杂度，使用map存储已经计算过的结果
func numTrees(num int, hashMap map[int]int) int {
	if num == 0 || num == 1 {
		return 1
	}
	count := 0
	if _, ok := hashMap[num]; ok{
		return hashMap[num]
	}
	for i:=1;i<=num;i++ {
		left := numTrees(i-1, hashMap)
		right := numTrees(num-i, hashMap)
		count += left * right
	}
	hashMap[num] = count
	return count
}

//动态规划，需要从将0,1,2....,num的节点全部计算出来，然后返回第num个节点
func numTrees2(num int) int {
	dp := make([]int, num+1)
	dp[0] = 1
	dp[1] = 1
	for i:=2;i<=num;i++ {
		for j:=1;j<=i;j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[num]
}

func main() {
	res := numTrees2(5)
	fmt.Println(res)
}
