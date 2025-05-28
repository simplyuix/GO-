// functions or types that work with any data type,

/* 
		With generics:

		You write the logic once.
		It works with many types.
		Compiler still checks types.

*/

package main
import (
	"fmt"
)

type N interface {
	int | float32
}
//  setting the type like deefine in C++ 


func add[ T N] (a,b T) T{
		return (a+b)
}

//  Generic don't allow + 

func main() { 
	
	v := add(10,5)
	fmt.Println("Value: ",v)
	
}