package main

import (
   "fmt"
)

//万能传参
func myFunc(arg interface{}){
	fmt.Println("myFunc is called")
	fmt.Println("arg = ",arg)
	//断言
	value,ok := arg.(string)
	if ok {
		fmt.Println("arg is string,value = ",value)
	}
}

type Test struct{
	name string
}

func main(){
   var test = Test{"zhangsan"}
   myFunc(test)
}
