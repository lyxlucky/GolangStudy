package main

import (
	"fmt"
)

//错误声明方式
//声明全局变量请使用方式一或者方式二
//gC := 123

func main(){
	//方式一
	var a int = 10
	fmt.Printf("a = %d, a type of %T",a,a)

	//方式二
	var test = "1234"
	fmt.Println()
	fmt.Printf("test = %s, test's type is %T",test,test)

	//方式三 有局限性，不能声明为全局变量
	power := 123
	fmt.Println()
	fmt.Printf("power = %d,power = %T",power,power)

	testString := "testString"
	fmt.Println()
	fmt.Printf("testString = %s,type of testString %T",testString,testString)

	//声明多个变量
	var xx,yy int = 100,200;
	fmt.Println("xx = ",xx,",yy = ",yy)
	var tt,gg = 100,"200";
	fmt.Println("tt = ",tt,"gg = ",gg)
}