package main

import "fmt"

/**
golang实现并发的方式主要有两种，锁和channel。
atomic 和 sync 包里的一些函数就可以对共享的资源进行加锁操作,atomic主要设计CAS和原子性操作，加锁主要是sync
是使用锁还是使用channel，https://blog.csdn.net/m0_43499523/article/details/86483484。
没有固定答案，优先使用channel
 */

/*
主goroutine中的mustCopy函数调用将返回， 然后调用conn.Close()
关闭读和写方向的网络连接。 关闭网络链接中的写方向的链接将导致server程序收到一个文件
（end-of-le） 结束的信号。 关闭网络链接中读方向的链接将导致后台goroutine的io.Copy函
数调用返回一个“read from closed connection”（“从关闭的链接读”）。
在后台goroutine返回之前， 它先打印一个日志信息， 然后向done对应的channel发送一个值。
主goroutine在退出前先等待从done对应的channel接收一个值。 因此， 总是可以在程序退出前
正确输出“done”消息。
顺序是 新起的goroutine->channel->主goroutine
*/
/*
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Printf("done")
		done <- struct{}{}
	}()
	mustCopy()
	conn.Close()
	<-done
}

func mustCopy() {

}
*/

/*
方法2，串联的channel，类似管道的使用
 */

func counter(out chan <- int)  {
	for x:=1;x<100;x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan <- int, in <- chan int)  {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <- chan int)  {
	for x := range in {
		fmt.Println(x)
	}
}

func main() {
	
	//natual := make(chan int)
	//squ := make(chan int)
	//
	//go func() {
	//	for x:=1;x<100;x++ {
	//		natual <- x
	//	}
	//	close(natual)  //不加的话会死锁
	//}()
	//
	//go func() {
	//	for x := range natual {
	//		//x := <- natual
	//		squ <- x * x
	//	}
	//	close(squ)  //不加会死锁
	//}()
	//
	//for x := range squ {
	//	fmt.Println(x)
	//}
	
	//进一步进行改进，使用函数进行提取
	natual := make(chan int)
	squa := make(chan int)
	go counter(natual)
	go squarer(squa, natual)
	printer(squa)
}
