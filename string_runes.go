package main


//   print output, scan input, and format strings

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	message := "Hello, \nGo!"
	message1 := "Hello, \tGo!"
	message2 := "Hello, Go!" 
	rawMessage := `Hello`

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	fmt.Println("Length of rawmessage variable is", len(rawMessage))

	fmt.Println("The first character in message var is", message[0]) 

	greeting := "Hello "
	name := "Alice"
	fmt.Println(greeting + name)

	str1 := "Apple"  
	str := "apple"  
	str2 := "banana" 
	str3 := "app"   
	fmt.Println(str1 < str2)
	fmt.Println(str3 < str1)
	fmt.Println(str > str1)
	fmt.Println(str > str3)

	for _, char := range message {
	
		fmt.Printf("%v\n", char)
	}

	fmt.Println("Rune count:", utf8.RuneCountInString(greeting))
	fmt.Println("Rune count:",len(greeting))

	greetingWithName := greeting + name
	fmt.Println(greetingWithName)

	var ch rune = 'a'
	jch := '日'

	fmt.Println(ch)
	fmt.Println(jch)

	fmt.Printf("%c\n", ch)
	fmt.Printf("%c\n", jch)

	cstr := string(ch)
	fmt.Println(cstr)

	fmt.Printf("Type of cstr is %T\n", cstr)

	const NIHONGO = "日本語" 
	fmt.Println(NIHONGO)

	jhello := "こんにちは" 
	for _, runeValue := range jhello {
		fmt.Printf("%c\n", runeValue)
	}

	r := '😊'
	fmt.Printf("%v\n", r)
	fmt.Printf("%c\n", r)
}