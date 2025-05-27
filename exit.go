package main

import (
	"fmt"
	"os"
)

func main() {

	

	fmt.Println("Starting the main function")
    defer fmt.Println("Deferred statement")
	
	os.Exit(404)

	
	fmt.Println("End of main function")

}