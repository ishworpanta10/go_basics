package main

// Concurrency : we can have multiple threads executing code. If one thread blocks, another one is picked up and worked on
// Parallelism : multiple threads executed at the exact same time. Requires multiple CPU's

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{"https://www.facebook.com", "https://www.google.com", "https://www.ishworpanta.com.np/?i=1", "https://www.amazon.com"}

	for _, link := range links {
		checkLink(link)
	}

}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		// fmt.Println(link, err)
		fmt.Println(link, "is down")
		return
	}

	fmt.Println(link, " is up and running")

}
