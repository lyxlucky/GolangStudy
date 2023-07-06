package main

import (
   "fmt"
)

func func1(){
	fmt.Println("A")
}

func func2(){
	fmt.Println("B")
}

func func3(){
	fmt.Println("C")
}

func returnFunc() int {
	fmt.Println("returnFunc -----")
	return 0;
}

func deferFunc(){
	fmt.Println("deferFunc -----")
}


func deferTest() int {
	defer deferFunc();
	return returnFunc();
}


func main(){ 
   fmt.Println("main start 1")
   //defer可以有多个，遵循栈模型
   //defer fmt.Println("main end")
   fmt.Println("main start 2")

   defer func1()
   defer func2()
   defer func3()
   //知识点：当defer和return在一起时，defer后执行，return 先执行

   e := deferTest()
   fmt.Println(e)
}
