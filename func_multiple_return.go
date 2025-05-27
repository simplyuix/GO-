package main

import (
	"errors"
	"fmt"
)

func main() { 
   
	r,q := help(10,6)
	if ( q==1 ) { 
		 fmt.Println(errors.New("Error!"))
	}

	fmt.Println(r)
	fmt.Println(q)

}

func help( a,b int)(r , q int) {
	r = a%b
	q = a/b
	return 
}

