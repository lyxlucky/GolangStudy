package main

import (
   "fmt"
)

//相当于int的一个别名
type myInt int;

//定义一个结构体
type Book struct{
	title string
	auth string
}

func changebook1(book Book){
	book.auth = "buzhidao"
}

func changebook2(book *Book){
	book.auth = "liao"
}


//值传递
func main(){
   var i myInt = 10
   fmt.Println("i = ",i)
   var book Book;
   book.title = "从入门到入土"
   book.auth = "Liao"
   changebook1(book)
   fmt.Println("book = ",book)
   changebook2(&book)
   fmt.Println("book = ",book)
}
