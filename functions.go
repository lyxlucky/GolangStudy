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