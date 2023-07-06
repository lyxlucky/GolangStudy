package main

import (
   "fmt"
)

func printMap(cityMap map[string]string){
	for i, v := range cityMap {
		fmt.Println("index = ",i,"value = ",v)
	}
}

func main(){
   cityMap := make(map[string]string)

   //添加
   cityMap["China"] = "Beijing"
   cityMap["Japan"] = "Tokyo"
   cityMap["KingDom"] = "Manchester"

   //修改
   cityMap["China"] = "ChangSha"

   //删除
   delete(cityMap,"Japan")

   //是一个引用传递
   printMap(cityMap)

}
