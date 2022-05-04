package main

import "fmt"

var res []string
func generateParenthesis(n int) []string {
	dfs(n, 0, 0, []byte{})
	return res
}

func dfs(n int, index int,right int, tmp []byte) {
	//cnt表示右括号的数量
	if index == 2*n {
		if right != n {
			return
		}
		res = append(res, string(tmp))
		return
	}
	//剪枝
	if right > len(tmp)-right || len(tmp)-right > n {
		return
	}
	dfs(n, index+1, right, append(tmp, '('))
	dfs(n, index+1, right+1, append(tmp, ')'))
}

func main() {
	n := 3
	a := generateParenthesis(n)
	fmt.Println(a)
}


