package main

import (
   "fmt"
)

func main(){
	c := make(chan int)
	go func(){
		defer fmt.Println("Goroutine结束")
		fmt.Println("Goroutine正在运行")
		//将666发送给c
		c <- 666
	}()

	num := <-c
	fmt.Println("num = ",num)
	fmt.Println("")
}
