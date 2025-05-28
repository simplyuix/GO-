package main

import (
	"fmt"
)

/*  
     Closure is something like { A function which return as a function } so, you can store the returned function ans use it many times
	 The main function is not instantiated evryting, until we can't call it directly 

	 functions for Maintaing states, without exposition the variable

    


*/

func main() { 
    a := adder()

	fmt.Println("Check the value: ", a())

	 a = adder()

	fmt.Println("Check the value: ", a())
	fmt.Println("Check the value: ", a())

	fmt.Println("________________________________________________________________________________")


	subtracter := func() func(int) int {

		countdown := 99
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	// Using the closure subtracter
	fmt.Println(subtracter(1))
	fmt.Println(subtracter(2))
	fmt.Println(subtracter(3))
	fmt.Println(subtracter(4))
	fmt.Println(subtracter(5))

}

// func adder() func() int {
// 	i := 0 
// 	fmt.Println("Previous value of i: ",i)
// 	return func() int { 
//            i++;
// 		   fmt.Println("Adding one to i: ", i);
// 		   return i
// 	}
// }

func adder() func() int {
	i := 0
	fmt.Println("previous value of i:", i)
	return func() int {
		i++
		fmt.Println("added 1 to i")
		return i
	}
}

/* 
		Previous value of i:  0
		Adding one to i:  1
		Check the value:  1
		Adding one to i:  2
		Check the value:  2
		Adding one to i:  3
		Check the value:  3

*/