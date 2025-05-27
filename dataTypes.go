package main

import (
	"fmt"
	// "go/types"
	"reflect"
)

//  reflect is used to get the type of function

func main() {


	fmt.Println("Hello" + "World")
	fmt.Println("9 X 10 =", 9*10)
	fmt.Println("180.18/2.0 = ", 180.18/2.0)
    
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	x := 34.4


	fmt.Println(reflect.TypeOf(x))
}