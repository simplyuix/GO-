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

}
