package main

import (
   "fmt"
)

//如果当前类大写，表示其他包也可以访问
type Hero struct{
	//如果说属性是大写，则表示该属性是公共属性，反之，则只能在本类中访问
	name string
	ad int
	levle int
}

//把当前方法绑定到结构体上
func (this Hero) getName(){
	fmt.Println("name = ",this.name)
}

//如果没有加*则表示是当前对象的一个副本
//并不会真正改变这个函数的值
func (this *Hero) setName(name string){
	this.name = name
}

func main(){
   hero := Hero{name:"liao",ad:100,levle:1}
   hero.getName();
   hero.setName("Yuxuan Liao")
   fmt.Println(hero)
}
