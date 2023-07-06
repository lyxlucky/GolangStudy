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
