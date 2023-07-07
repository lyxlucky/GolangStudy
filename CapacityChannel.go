package main

import (
   "fmt"
   "time"
)

func main(){
   
	//开辟一个channel 数据类型为int 容积为3
	c := make(chan int,3)

	fmt.Println("len(c) = ",len(c),"cap(c) = ",cap(c))

	go func(){
		defer fmt.Println("子Go程结束")
		for i := 0; i < 4; i++ {
			//把值赋给Channel 
			c <- i
			fmt.Println("子Go程正在运行，发送的元素 ",i,"len(c) = ",len(c),"cap(c) = ",cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <- c
		fmt.Println("num = ",num)
	}
	defer fmt.Println("main程结束")
}
