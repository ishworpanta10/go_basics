package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	fmt.Println("========= Interface Http ==============")
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos")

	if err != nil {
		fmt.Println("Error in Get", err)
		os.Exit(1)
	}

	// bs := make([]byte, 9999)
	// response.Body.Read(bs)
	// fmt.Println("Response Body from Get : ", string(bs))

	// io.Copy(os.Stdout, response.Body)

	lg := logWriter{}
	io.Copy(lg, response.Body)

}

func (logWriter) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}
