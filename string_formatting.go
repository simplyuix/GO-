package intermediate

import "fmt"

func main() {

	num := 424

	// means length shoudl be 5 if not add 0

	fmt.Printf("%05d\n", num)

	message := "Hello"

	// length 10 if not add space

	fmt.Printf("|%10s|\n", message)
	fmt.Printf("|%-10s|\n", message)

	message1 := "Hello \nWorld!"
	message2 := `Hello \nWorld!`  // raw string wil print as it is

	fmt.Println(message1)
	fmt.Println(message2)


}