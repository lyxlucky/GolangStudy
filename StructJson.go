package main

import (
   "fmt"
   "encoding/json"
)

type Movie struct{
	Title string
	Year int
	Price int
	Actors []string
}

func main(){
   //将结构体----》转成json
   movie := Movie{"喜剧之王",2000,10,[]string{"zhangsan","lisi"}}
   jsonStr,_ := json.Marshal(movie)
   fmt.Printf("json = %s",jsonStr)

   //将json解析为结构体
   m := Movie{}
   json.Unmarshal(jsonStr,&m)
   fmt.Println("m = ",m)

}
