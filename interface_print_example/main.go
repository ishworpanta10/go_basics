package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func main() {

	fmt.Println("====== Interface with Local file read ========")

	fmt.Println("Reading text file named :", os.Args[1])

	fileData, err := os.Open(string(os.Args[1]))
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)

	}

	lg := logWriter{}
	// io.Copy(os.Stdout, fileData) // OR custom log with interface
	io.Copy(lg, fileData)

}

func (logWriter) Write(bs []byte) (n int, err error) {

	fmt.Println("========== Printing Content of file ===========")
	fmt.Println(string(bs))
	fmt.Println("Writen byte of length :", len(bs))
	return len(bs), nil
}
