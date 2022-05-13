package main

import "fmt"

//空间复杂度为1的算法，自我翻转
func rotate(matrix [][]int)  {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	n := len(matrix)
	for i:=0;i<n;i++ {
		for j:=i+1;j<n;j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i:=0;i<n;i++ {
		for j:=0;j<n/2;j++ {
			matrix[i][j], matrix[i][n-i-1] = matrix[i][n-i-1], matrix[i][j]
		}
	}
}

func main()  {
	matrix := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	rotate(matrix)
	fmt.Println(matrix)
}
