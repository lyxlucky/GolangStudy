package main

import "fmt"

//const 定义枚举类型
const (
	//可以在const ()里添加一个关键字，iota，每行的iota会默认加一，默认值为0;
	BEIJING = 10 * iota
	SHANGHAI
	SHENZHEN
)

//注意！！！
//这里的iota是每行加一，每行！！！
//iota只能配合const使用！！
const (
	a,b = 1 + iota , 2 + iota //a = 1 b = 2
	c,d  					  //b = 2 b = 3
	e,f						  //e = 3 f = 4
)



func main(){
	//常量声明方式（只读）
	const length = 10
	fmt.Println("length = ",length)
	fmt.Println("beijing = ",BEIJING,"shanghai = ",SHANGHAI,"shenzhen = ",SHENZHEN)
	fmt.Println("-------------------------------------------")
	fmt.Println("a = ",a,"b = ",b,"c = ",c,"d = ",d,"e = ",e,"f = ",f)
}