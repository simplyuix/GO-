/* 
       How to take input from user : 

			fmt.Scanln { to take single input }
			Scanln(&N) { we need to provide address where to store }
		
	  bufio stands for buffered I/O (input/output).


	  In order to take entire line as Input : 

	  		reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')

*/


/*  Printing sum to 2 numbers */


		// package main
		// import (
		// 	"fmt" 
		// 	//  "bufio"
		// 	)

		// 	func main() { 

		// 	var a,b int64
		// 	fmt.Scanln(&a,&b)
		// 	fmt.Println(a+b)

		// }
	
/*    Basics   */

package Basics

import (
	"fmt"
	// "math"
)

func main() { 

	var a,b int64
    a,b = 20,6
	
	result := a+b 
	fmt.Println("Sum of two",result)

	result = a-b
	fmt.Println("Diff of two numbers", result)

	result = a/b 
    fmt.Println("Division of two number", result)

	result = a*b
	fmt.Println("Multiplication of two number",result)
}




