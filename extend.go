package main

import (
   "fmt"
)

//父类
type Human struct{
	name string
	age string
}

func (this *Human) getName() string {
	return this.name;
}

func (this *Human) setName(name string){
	this.name = name
}

func (this *Human) eat(){
	fmt.Println("human is eat()-----")
}


//子类
type SuperMan struct{
	Human //SuperMan包含了human类的属性	
	level int
}

//重写
func (this *SuperMan) eat(){
	fmt.Println("Superman is eat()---")
}

//子类新方法
func (this *SuperMan) fly(){
	fmt.Println("superman is flying()----")
}


func main(){
   s := Human{"zhangsan","18"}
   name := s.getName()
   fmt.Println(name)
   s.eat()

   r := SuperMan{Human{"zhangsan","12"},1}
   r.eat();
   fmt.Println(r.getName())
}
