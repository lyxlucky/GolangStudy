package main

//导包需在gopath下
import (
	"fmt"
	"lib1"
	//匿名导包，不想使用lib1中的方法，但需使用lib1的init()函数
	//_"lib1"
	// 起别名，可以通过别名来调用库中的方法。
	// mylib2 "lib1"
	//将lib1中的所有方法导入本作用域，不在需要使用库名.方法这种方式。
	//."lib1"
)

func main(){
	//默认会去执行lib1包下面的初始化方法（有点类似于java的构造函数）
	lib1.lib1Test()
}