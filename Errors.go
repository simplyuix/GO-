package main

import (
	"errors"
	"fmt"
)


func canVote(age int) error {
	if age < 18 {
		return errors.New("you are not eligible to vote")
	}
	return nil
}
var ErrUnauthorized = errors.New("unauthorized")

func login() error {
	return ErrUnauthorized
}
type customError struct {
	code    int
	message string
	er      error
}

func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s, %v\n", e.code, e.message, e.er)
}


func doSomething() error {
	err := doSomethingElse()
	if err != nil {
		return &customError{
			code:    500,
			message: "Something went wrong",
			er:      err,
		}
	}
	return nil
}

func doSomethingElse() error {
	return errors.New("internal error")
}

func main() {
	
	err := canVote(16)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("You can vote!")
	}

	
	err = canVote(20)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("You can vote!")
	}

	err = login()
	if errors.Is(err, ErrUnauthorized) {
		fmt.Println("Unauthorized error occurred")
	}

	err = doSomething()
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("Operation completed successfully!")

}
