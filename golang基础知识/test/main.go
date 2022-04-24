package main

import "fmt"

func add()  {
	var a int
	var k = 0
NextAdd:
	for k < 5 {
		for i:=0;i<100;i++ {
			if i==3 {
				continue
			}
			a = a + i
		}
		k++
		fmt.Println("k=", k)
		fmt.Println("a=", a)
		continue NextAdd
	}
}

func main() {
	add()
}


