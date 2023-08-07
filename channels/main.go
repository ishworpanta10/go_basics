package main

// Concurrency : we can have multiple threads executing code. If one thread blocks, another one is picked up and worked on
// Parallelism : multiple threads executed at the exact same time. Requires multiple CPU's

// Go Routine:
// - program start with a single main routine
// - go keyword is used to create other child routine
// - main routine does not care about child routine to complete
// - so we introduce CHANNEL

// Channel:
// - channel are the intermediate communication between different go routines (main,child,child)

// Sending data with channels
// - channel <- 5  => send value 5 to this channel
// - myNumber <- channel => wait for a value to be send to channel, when we get that value then assign it to myNumber
// - fmt.Println(<- channel) => wait for a value to be send to channel, when we get log it immediately

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"https://www.facebook.com",
		"https://www.google.com",
		// "https://www.ishworpanta.com.np/?i=1",
		"https://www.amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	fmt.Println(<-c)

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		// fmt.Println(link, err)
		c <- "Might be down"
		fmt.Println(link, "is down")
		return
	}

	fmt.Println(link, " is up and running")
	c <- "is runnning"

}
