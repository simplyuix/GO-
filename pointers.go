package main

import (
	"fmt"
)


//  & gives address of the variable


func main() { 

	var a int = 10 

	var b *int = &a
	fmt.Println("Value: ",a) ;

	fmt.Println("Address: ",&a) ;

	fmt.Println("Address of B : ", &b) ;
	*b++
    (*b)++
	fmt.Println("Value of B : ", b) ; // Dereferrencing 
	fmt.Println("Value: ",a) ;



}