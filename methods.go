package main

import (
	"fmt"
)


type Persona struct { 
	FirstName string
	LastName string
	age  int64
}

func main() { 

	/*   
	    Difference between method and function: 

		  both are same but methods are associated with types like structs : 
		  
		  func( receiver ) name(...) { 
		  
		  }
	*/


p := Persona {
	FirstName: "Aman",
	LastName: "Singh",
	age: 19,
}
fmt.Println(p.fullName())





}

func (x Persona)fullName()string{
	return x.FirstName+" "+x.LastName

}

/* 
     Struct and methods are to be defined outside of main function  
	 Go doesn't allow to define in main function 
    

*/