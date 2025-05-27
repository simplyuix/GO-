package main

import "fmt"

func sub(a,b int)int{
		return (b-a)
	}
func dummyFunction (a,b int , operation func(int,int) int ) int {
            return operation(a,b)
	}
func main() {

	

	// sum := add(1, 2)
	fmt.Println(add(2, 3))

	greet := func() {
		fmt.Println("Hello Anonymous Function")
	}

	greet()

	sayHi := func(){
		fmt.Println("Hello, Hi Namaskar!")
	}

	sayHi();
    

	test := dummyFunction(5,6,sub)
	
	
	fmt.Println("value : ", test)


	result := applyOperation(5, 3, add)
	fmt.Println("5 + 3 =", result)

	
	multiplyBy2 := createMultiplier(2)
	fmt.Println("6 * 2 = ", multiplyBy2(6))

}

func add(a, b int) int {
	return a + b
}

func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}