package main
import (
	"fmt" 
	GET "net/http"
)

func main() { 
	 json,err :=  GET.Get("https://jsonplaceholder.typicode.com/todos")
	if(err!=nil) {
		fmt.Println("Error! ",err)
	} else {
		fmt.Println("Json ",json)
	}

	defer json.Body.Close()
}