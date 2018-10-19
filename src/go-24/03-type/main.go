package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var s string = "Subhransu"
	var sV string = "true"
	// var i int = 10
	// var f float32 = 1.2

	fmt.Println(reflect.TypeOf(s))

	b, err := strconv.ParseBool(sV)

	if err == nil {
		fmt.Println(reflect.TypeOf(b))
	}
	
}
