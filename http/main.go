package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("========= Interface Http ==============")
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos")

	if err != nil {
		fmt.Println("Error in Get", err)
	}

	fmt.Println("Response from Get : ", response)
}
