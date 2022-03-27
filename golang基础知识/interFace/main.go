package main

// 接口本质上就是做一个约定，在接口中定义好约定的方法，你使用我这个接口就必须要按照这种约定实现接口的所有方法
//比如sort接口
/*
sort包内置的提供了根据一些排序函数来对任何序列排序的功能。 它的设计非常独
到。 在很多语言中， 排序算法都是和序列数据类型关联， 同时排序函数和具体类型元素关
联。 相比之下， Go语言的sort.Sort函数不会对具体的序列和它的元素做任何假设。 相反， 它
使用了一个接口类型sort.Interface来指定通用的排序算法和可能被排序到的序列类型之间的约
定。 这个接口的实现由序列的具体表示和它希望排序的元素决定， 序列的表示经常是一个切
片。
 */
// 或者可以说接口是实现多态的一种方式

type Interface interface {
	Len() int
	Less(i, j int) bool // i, j are indices of sequence elements
	Swap(i, j int)
}

func main() {

}

