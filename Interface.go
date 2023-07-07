package main

import (
   "fmt"
)

type Animal interface{
	sleep()
	getColor() string
	getType() string
}

type Cat struct{
	color string
}

func (this *Cat) sleep(){
	fmt.Println("cat is sleep")
}

func (this *Cat) getColor() string {
	return this.color
}

func (this *Cat) getType() string {
	return "dog"
}

func main(){
   var animal Animal;
   animal = &Cat{"绿色"}
   animal.sleep()
}
