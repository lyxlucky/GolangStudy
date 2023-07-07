package main

import (
   "fmt"
)

func main(){
   

	c := make(chan int,5)

	go func(){
		for i := 0; i < 5; i++ {
			c <- i
		}
		defer close(c)
	}()

	//可以使用range来不断操作channel
	for data := range c{
		fmt.Println(data)
	}
}
