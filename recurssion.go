package main

import (
	"fmt"
)

func main() { 


	fmt.Println("Let's Print Sum of First n numbers")
	var N int64 
	fmt.Scanln(&N)
	val := sum(N)
	fmt.Println(val)

	fmt.Println("Let's Print fact of First n numbers")
	fmt.Scanln(&N)
	val = fact(N)
	fmt.Println(val)




}
func sum(N int64) (int64) {
    if(N==0) {return 0 
	} else {
		return N+sum(N-1)
	}
}


func fact(N int64) (int64) { 

	if(N==0) {
		return 1
	} else {
		return N*fact(N-1)
	}

}