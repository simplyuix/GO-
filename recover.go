package main

import (
	
	"fmt"
)

func main() { 
//    recover only work inside a defer function 

defer func() {
		
		r := recover()
		if r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	fmt.Println("Start Process")
	// panic("Something went wrong!")
	fmt.Println("End Process")

}