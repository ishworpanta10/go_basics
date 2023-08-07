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

//Note:
//- when main routine get value from channel it immediatley exit program as it has no other task to do

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"https://www.facebook.com",
		"https://www.google.com",
		"https://www.ishworpanta.com.np/?i=1",
		"https://www.amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c) => this will make main routine wait and program does not stop and hang, we have 4 slice data but print 5

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	/// Repeated Go Routine channel

	// for {
	// 	go checkLink(<-c, c)
	// }

	//same syntax
	for l := range c {

		go func(link string) {
			time.Sleep(4 * time.Second)
			go checkLink(link, c)
		}(l)

	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		// fmt.Println(link, err)
		// c <- "CHANNEL MESSAGE : MIGHT DOWN"
		c <- link
		fmt.Println(link, "might down")
		return
	}

	fmt.Println(link, " is up and running")
	// c <- "CHANNEL MESSAGE : RUNNING"
	c <- link

}
