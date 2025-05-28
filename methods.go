package main

import (
	"fmt"
)

type Persona struct {
	FirstName string
	LastName  string
	age       int64
}

func main() {

	/*
		    Difference between method and function:

			  both are same but methods are associated with types like structs :

			  func( receiver ) name(...) {

			  }


			  If we use a pointer receiver ( for updating the value of struct )
	*/

	p := Persona{
		FirstName: "Aman",
		LastName:  "Singh",
		age:       19,
	}
	fmt.Println(p.fullName())

	k := Person{Name: "Aman", Age: 20}

	k.Greet()       // Output: Hello, Aman
	k.Birthday()    // Age becomes 21
	fmt.Println(k)  // {Aman 21}

}

func (x Persona) fullName() string {
	return x.FirstName + " " + x.LastName

}

type Person struct {
	Name string
	Age  int
}

// Pointer receiver
func (p *Person) Birthday() {
	p.Age++
}

// Value receiver
func (p Person) Greet() {
	fmt.Println("Hello,", p.Name)
}


/*
     Struct and methods are to be defined outside of main function
	 Go doesn't allow to define in main function


*/
