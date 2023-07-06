# Golang学习



### Golang背景（摘自wikipedia）

Go（又称Golang[4]）是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。

罗伯特·格瑞史莫、罗勃·派克及肯·汤普逊于2007年9月开始设计Go，[3]稍后伊恩·兰斯·泰勒（Ian Lance Taylor）、拉斯·考克斯（Russ Cox）加入项目。Go是基于Inferno操作系统所开发的。[5]Go于2009年11月正式宣布推出，成为开放源代码项目，支持Linux、macOS、Windows等操作系统

### Golang描述（摘自Wikipedia）

Go的语法接近C语言，但对于变量的声明有所不同。Go支持垃圾回收功能。Go的并行计算模型是以东尼·霍尔的通信顺序进程（CSP）为基础，采取类似模型的其他语言包括Occam和Limbo[3]，Go也具有这个模型的特征，比如通道传输。通过goroutine和通道等并行构造可以建造线程池和管道等[8]。在1.8版本中开放插件（Plugin）的支持，这意味着现在能从Go中动态加载部分函数。

与C++相比，Go并不包括如枚举、异常处理、继承、泛型（此功能在go1.18中加入）、断言、虚函数等功能，但增加了切片(Slice) 型、并发、管道、垃圾回收、接口等特性的语言级支持[3]。对于断言的存在，则持负面态度，同时也为自己不提供类型继承来辩护。

不同于Java，Go原生提供了关联数组（也称为哈希表（Hashes）或字典（Dictionaries））

#### Golang HelloWorld

```go
package main

/**
 * 
 import(
	"fmt"
	"time"
 ) 
 */
import "fmt"
import "time"

func main() {
	time.Sleep(1000)
	fmt.Println("hello world")
}
```

注意点：

- 在Golang中，括号不能单独起一格，必须按照上述代码格式（语法规定，否则会报错）

#### Golang变量声明方式

```go
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
```

golang大致有四种变量声明方式

- 第一种 `var a int = 0` 这种属于显示声明
- 第二种 `var a = 0` 这种属于隐式声明
- 第三种 `a : =  0`，这种方法最为常用，也有点类似Java语法糖的意思（个人见解）**注意，这种方法不可用于声明全局变量，声明全局变量请采用其他方式**
- 第四种 `var a,b = 0,0`多变量声明方式

#### Golang常量

```go
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
```

在Golang中使用`const`关键字来定义常量，同时可以配合关键字`iota`来使用，详情请见以上代码。

#### Golang函数

```go
package main

import "fmt"

/*
这个方法表示两个形参，返回值为int类型
*/
func foo1(a string,b int) int {
	fmt.Println("-----foo1-----")
	fmt.Println("a == ",a,"b == ",b)
	c := 100
	return c;
}


/*
表示返回两个参数，都是整形的（匿名）
*/
func foo2(a string,b int) (int,int) {
	fmt.Println("-----foo2-----")
	fmt.Println("a =",a,"b = ",b)
	return 100,200;
}

/*
返回多个形参，有形参名称不匿名！！
*/
func foo3(a string,b int) (r1 int,r2 int) {
	fmt.Println("-----foo3-----")
	fmt.Println("a =",a,"b = ",b)
	r1 = 300
	r2 = 400

	// 形参一旦初始化，默认值为0
	//r1 r2 的作用域为foo3

	return
}

/*
函数第四种声明形式(多返回值)
*/
func foo4(a string,b int) (r1,r2 int) {
	fmt.Println("-----foo4-----")
	fmt.Println("a =",a,"b = ",b)
	r1 = 300
	r2 = 400

	return
}

func main(){
	c := foo1("123",1)
	fmt.Println("foo1's c == ",c)

	f1,f2 := foo2("321",2)
	fmt.Println("foo2's f1 == ",f1,"foo2 's f2 == ",f2)

	f1,f2 = foo3("asd",3)
	fmt.Println("foo3 's f1 = ",f1,"foo3 's f2 = ",f2)

	f1,f2 = foo4("asd",3)
	fmt.Println("foo4 's f1 = ",f1,"foo3 's f2 = ",f2)
}
```

在golang中函数的声明形式大致可以分为以上四种，**在Golang中函数不支持重载**

#### Golang导包

```go
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
```

在Golang中，导包默认会去导入Gopath下的包，导入一个包会默认执行init()函数，同时我们也可以匿名导包（只执行包下的init函数），具体请参照上面代码

#### Golang指针

在Golang中也有指针的概念，但是用的比较少

```go
package main

import (
   "fmt"
)

func changeValue(a *int){
	*a = 10
}

//交换两个值
func swap(a *int,b *int){
	var temp int = *a
	*a = *b
	*b = temp;
}


func main(){
	a := 1
	changeValue(&a)
	fmt.Println("a = ",a)
	var w,s int = 1,2
	swap(&w,&s)
	fmt.Println("w = ",w,"s = ",s)

	var p *int
	p = &a
	fmt.Println(p)
	fmt.Println(&a)
}
```

#### Golang关键字defer

defer关键字，用于在函数执行完毕后执行（类似于java的finally），**defer遵循栈模型，压栈，先进后出**

```go
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
```

#### Golang数组

在golang中也分静态数组和动态数组，静态数组严格匹配数组类型，而动态数组则不会（更推荐使用动态数组）

静态数组和动态数组的区别:

- 静态数组严格匹配数组类型
- 静态数组是值传递
- 动态数组是引用传递

静态数组：

```go
package main

import (
   "fmt"
)

//严格匹配数组类型 (这里是3就只能传长度为3的数字)
func printArray(array [3]int){
	//值拷贝
	for index,value := range array {
		fmt.Println("index = ",index,"value = ",value)
	}
}


func main(){
	
	//静态数组
	var myArray [10]int

	myArray2 := [3]int{1,2,3}

	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}

	for index, value := range myArray2 {
		fmt.Println("index = ",index,"value = ",value)
	}

	printArray(myArray2)

}	
```



动态数组

```go
package main

import (
   "fmt"
)

//动态数组是引用传递 不同元素长度的动态数组他们的形参是一致的
func printArray(array []int){
    //"_"表示匿名
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
```

