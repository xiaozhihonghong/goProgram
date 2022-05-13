package main

import (
	"fmt"
	"sort"
	"strings"
)

//这题的关键是使用一个map，map的key是排序后的字符串，value是原字符串
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}
	mp := make(map[string][]string)
	for i:=0;i<len(strs);i++ {
		sortStr := sortString(strs[i])
		mp[sortStr] = append(mp[sortStr], strs[i])
	}
	res := make([][]string, 0)
	for _, value := range mp {
		res = append(res, value)
	}
	return res
}

func sortString(str string) string {
	split := strings.Split(str, "")
	sort.Strings(split)
	return strings.Join(split, "")
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(strs)
	fmt.Println(res)
}
