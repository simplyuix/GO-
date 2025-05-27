package main

import (
	// "errors"
	"fmt"
)

func main() { 
   
	fmt.Println(divide(10,2))

}

func divide(a, b int) int {
    if b == 0 {
        panic("cannot divide by zero")
    }
    return a / b
}
//  exists from the execution directly
