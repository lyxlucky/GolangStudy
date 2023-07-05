package main

import "fmt"

//const 定义枚举类型
const (
	//可以在const ()里添加一个关键字，iota，每行的iota会默认加一，默认值为0;
	BEIJING = 10 * iota
	SHANGHAI
	SHENZHEN
)

func main(){
	//常量声明方式（只读）
	const length = 10
	fmt.Println("length = ",length)
	fmt.Println("beijing = ",BEIJING,"shanghai = ",SHANGHAI,"shenzhen = ",SHENZHEN)
}