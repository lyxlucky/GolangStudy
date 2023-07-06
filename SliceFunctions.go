package main

import (
   "fmt"
)

/*
长度表示左右指针的距离
切片的扩容机制，当长度超过容积时，会扩容至原来的两倍
*/
func main(){
	// 定义一个长度为3但是体积为5的slice切片
   var slice = make([]int,3,5)
   fmt.Println("slice = ",len(slice),"slice cap = ",cap(slice))

   slice = append(slice,1)
   fmt.Println("slice = ",len(slice),"slice cap = ",cap(slice))

   slice = append(slice,1)
   fmt.Println("slice = ",len(slice),"slice cap = ",cap(slice))

   //切片的截取
   num := []int{1,2,3,4,5}
   split := num[1:]
   fmt.Println("split = ",split)

   //切片的copy
   copy(slice,num)
   fmt.Println("slice = ",slice,"slice cap = ",cap(slice))

}
