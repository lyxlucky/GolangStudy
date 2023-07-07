package main

import (
   "fmt"
   "reflect"
)

type resume struct{
	name string `info:name doc:我的名字`
	sex string `info:sex`
}

func findTag(str interface{}){
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		info := t.Field(i).Tag.Get("info")
		doc := t.Field(i).Tag.Get("doc")
		fmt.Println("info = ",info,"doc = ",doc)
	}
}

func main(){
   var re resume
   findTag(&re)
}
