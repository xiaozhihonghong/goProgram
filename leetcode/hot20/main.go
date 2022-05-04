package main

import "fmt"

//有效的括号

func isValid(s string) bool {
	if len(s)%2==1 || len(s)==0 {
		return false
	}
	stack := make([]byte, 0)
	for i:=0;i<len(s);i++ {
		switch s[i] {
		case ')':
			ch := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if ch != '(' {
				return false
			}
		case ']':
			ch := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if ch != '[' {
				return false
			}
		case '}':
			ch := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if ch != '{' {
				return false
			}
		default:
			stack = append(stack, s[i])
		}
	}
	return true
}

func main() {
	s := "(}"
	res := isValid(s)
	fmt.Println(res)
}
