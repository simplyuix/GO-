package main

import (
	"bufio"
	"fmt"
	"os"
	"text/template"
)

type Name struct {
	Name string 
}

func main() {
	t := template.New("Check")

	
	t, err := t.Parse("Welcome! {{.Name}}")
	if err != nil {
		fmt.Println("Parse error:", err)
		return
	}

	p := Name{
		Name: "Aman", 
	}

	err = t.Execute(os.Stdout, p)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("\nNo error")
	}

	tmpl := template.New("example")

	tmpl, err1 := template.New("example").Parse("Welcome, {{.name}}! How are you doing?\n")
	if err1 != nil {
		panic(err1)
	}
	tmpl = template.Must(template.New("example").Parse("Welcome, {{.name}}! How are you doing?\n"))

	
	data := map[string]interface{}{
		"name": "Aman",
	}

	err1 = tmpl.Execute(os.Stdout, data)
	if err1 != nil {
		panic(err1)
	}


	var N string 
	fmt.Scanf("%10s",&N)
	fmt.Println("Hello! ", N)
  
	reader := bufio.NewReader(os.Stdin) 
	fmt.Println("Provide Name")
	a,b,c := reader.ReadLine();

	fmt.Println("Hello!",string(a))
	fmt.Println("b! ",b)
	fmt.Println("c! ",c)
	




}
