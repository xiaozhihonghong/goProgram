//package main
//
//import "fmt"
//
//func main() {
//	arr := [8]int{}
//	for i := 0; i < 8; i++ {
//		arr[i] = i
//	}
//
//	fmt.Println("arr1= ", arr)
//	exchange(arr)
//	fmt.Println("arr2= ", arr)
//}
//
//func exchange(arr [8]int) {
//	for k, v := range arr {
//		arr[k] = v * 2
//	}
//}

//package main
//
//import "fmt"
//
//func main() {
//	arr := [8]int{}
//	for i := 0; i < 8; i++ {
//		arr[i] = i
//	}
//
//	fmt.Println("arr1= ", arr)
//	exchangeByAddress(&arr)
//	fmt.Println("arr2= ", arr)
//}
//
//func exchangeByAddress(arr *[8]int) {
//	for k, v := range *arr {
//		arr[k] = v * 2
//	}
//}

package main

import "fmt"

func main() {
	slice := []int{1,2,3,4,5}
	fmt.Println(slice)
	exchangeSlice(slice)
	fmt.Println(slice)
}

func exchangeSlice(slice []int) {
	for k, v := range slice {
		slice[k] = v * 2
	}
}
