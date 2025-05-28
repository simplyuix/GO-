package main
import (
	"fmt"
)
func main() { 
	/* 
	    %v Default Value
		%T Type of variable

		%s Plain String
		%q Double-Qouted String

		` ` -> in back tick String literal

	*/

	 i := 10 

	 fmt.Println(fmt.Printf("%v",i))
	 fmt.Println(fmt.Printf("%#v",i))

	 type Person struct { Name string }
		p := Person{"Aman"}
		fmt.Printf("%v\n", p)   
		fmt.Printf("%#v\n", p) 
	
	   fmt.Println(fmt.Printf("%T\n",i))

	   fmt.Printf("%T\n",i)

	   f := 10.3

	   fmt.Printf("%d\n",f)

	   sqlQuery := `SELECT * FROM users WHERE age > 30`

	   fmt.Printf("%s",sqlQuery)




}