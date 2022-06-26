package 切片

//切片是作为参数传入函数中是应用传递还是值传递，扩容时又会有什么现象？
//https://www.jianshu.com/p/7779ea76e1e9

//是值传递，作为参数的时候会进行切片拷贝，但指针指向的数据是相同的，所以形参影响实参，
//但是如果切片进行扩容时，拷贝成新的数组，那么指针指向不同的数据，形参就不影响实参了
