package main

import (
   "fmt"
)

//动态数组是引用传递 不同元素长度的动态数组他们的形参是一致的
func printArray(array []int){
	for _, v := range array {
		fmt.Println("v = ",v)
	}
	array[0] = 100
}

func main(){
   array := []int{123,321,1234}
   printArray(array)
   fmt.Println("------------------------------------")
   for _, v := range array {
	fmt.Println("v = ",v)
   }

}
