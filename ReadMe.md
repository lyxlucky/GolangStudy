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