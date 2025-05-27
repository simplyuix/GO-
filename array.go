package main

import (
	"fmt"
	// "go/types"
)

func main() { 


	//  var arr = [size]type[]

	 var arr [5]int
	 fmt.Println(arr) // [0 0 0 0 0]

	 arr[4] = 20
	 fmt.Println(arr) // [0 0 0 0 20]

	 for i:= 0 ; i<len(arr) ; i++ { 
		fmt.Println(arr[i])
	 }


	 for index,value := range arr { 
		fmt.Printf("Index %d: value %d: ",index,value)
	 }

	//   matrix 

	var matrix [3][3] int = [3][3]int{
			{0,1,0},
			{1,0,0},
			{0,0,1},

	 }

	  for index,value := range matrix { 
		fmt.Printf("Index %d: value %d: ",index,value)
	 }


}