package main

import (
   "fmt"
   "time"
   "runtime"
)


//子Goroutine
func task(){
	for i := 0;; i++ {
		fmt.Println("Hello Goroutine Go Thread")
		time.Sleep(1 * time.Second)
	}
}

//主Goroutine
func main(){
    //go task()
	// for i := 0;; i++ {
	// 	fmt.Println("Hello Goroutine main Thread")
	// 	time.Sleep(1 * time.Second)
	// }

	//匿名协程
	go func(){
		defer fmt.Println("A.defer")

		fmt.Println("A")
		//Go程退出
		runtime.Goexit()
	}()

	//匿名方法传参方式
	//这里拿不到返回值，必须使用Channel来实现线程之间的通信
	go func(a int,b int) bool {
		fmt.Println("a = ",a,"b = ",b)
		return true
	}(10,20)

	for i := 0;; i++ {
		time.Sleep(1 * time.Second)
	}
	fmt.Println("main exit")
}
