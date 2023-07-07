package main

import (
   "fmt"
)

//关闭channle之后还是能接受数据的
//对于nil channel 无论收发都会阻塞
func main(){
	//无缓冲区channel
	c := make(chan int)

	go func(){
		for i := 0; i < 5; i++ {
			c <- i
		}
		//关闭channel
		close(c)
	}()

	for{
		if data,ok := <- c;ok{
			fmt.Println(data)
		}else{
			break
		}
	}
	fmt.Println("Main Finished")

}
