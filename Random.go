package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	fmt.Println(rand.Intn(101))

	valr := rand.New(rand.NewSource(time.Now().Unix()))

	fmt.Println(valr.Intn(45))
}