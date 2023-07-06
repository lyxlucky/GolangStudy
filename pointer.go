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
