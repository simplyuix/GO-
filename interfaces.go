/* 
   Type that implements the required methods satisfies the interface.
*/

package main

import (
	"fmt"
	"math"
)


type Shape interface {
	Area() float64
}


type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}


func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}


//  this check the method implemented has the method implemented 

func printArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 3}

	printArea(c) 
	printArea(r) 
}
