package main

import (
   "fmt"
)

func main(){
   //声明slice是一个切片，默认值为1，2，3，长度为3
   slice1 := []int{1,2,3}
   var slice []int = make([]int,3)
   fmt.Println("slice length",len(slice1))
   fmt.Println("slice = ",slice)

   //判断一个slice是否为空
   if slice == nil{
	fmt.Println("slice为空")
   } else {
	fmt.Println("slice不为空")
   }
}
