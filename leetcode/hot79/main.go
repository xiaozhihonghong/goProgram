package main

import "fmt"

func exist(nums [][]byte, word string) bool {
	vis := make([][]bool, len(nums))
	for i:=0;i<len(vis);i++ {
		vis[i] = make([]bool, len(nums[0]))
	}
	for i:=0;i<len(nums);i++ {
		for j:=0;j<len(nums[0]);j++ {
			if dfs(nums, word, "", i, j, 0, vis) == true {
				return true
			}
		}
	}
	return false
}

// 可以优化，不需要s， index==len(word)就可以作为退出条件
func dfs(nums [][]byte, word string, s string, x, y int, count int, vis [][]bool) bool {
	if s == word {
		return true
	}
	if x < 0 || x >= len(nums) || y < 0 || y >= len(nums[0]) || nums[x][y] != word[count] || vis[x][y] {
		return false
	}
	vis[x][y] = true
	if dfs(nums, word, s+string(nums[x][y]), x+1, y, count+1, vis) ||
	dfs(nums, word, s+string(nums[x][y]), x-1, y, count+1, vis) ||
	dfs(nums, word, s+string(nums[x][y]), x, y+1, count+1, vis) ||
	dfs(nums, word, s+string(nums[x][y]), x, y-1, count+1, vis) {
		return true
	}
	vis[x][y] = false
	return false
}

func main() {
	board := [][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}
	word := "ABCCED"
	fmt.Println(exist(board, word))
}
