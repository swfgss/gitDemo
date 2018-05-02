package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int64  `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
	Age  int64  `json:"age" bson:"age"`
}

func main() {
	u := &User{10000, "jasonlin", 22}
	elems := reflect.TypeOf(u).Elem()
	for i := 0; i < elems.NumField(); i++ {
		fmt.Println(elems.Field(i).Tag.Get("json"))
	}
}
