package main

import (
   "fmt"
)

func main(){
	//定义一个键为stringvalue为string的map
   var myMap map[string]string
   if myMap == nil {
	fmt.Println("map为空")
   }
   //开辟空间
   myMap = make(map[string]string,10)

   myMap["one"] = "java"
   myMap["two"] = "python"
   myMap["three"] = "js"
   fmt.Println(myMap)

   //第二种声明方式
   myMap2 := make(map[int]string,3)
   myMap2[1] = "Java"
   myMap2[2] = "js"
   myMap2[3] = "python"

   fmt.Println(myMap2)

   //第三种声明方式
   mymap3 := map[string]string{
		"one" : "Java",
		"two" : "python",
		"three" : "js",
   }
   fmt.Println(mymap3)
}
