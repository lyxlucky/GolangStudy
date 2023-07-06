package lib

import (
	"fmt"
)

//当前lib一提供的接口
func lib1Test(){
	fmt.Println("lib1Test()")
}

func init(){
	fmt.Println("lib1 is init()")
}