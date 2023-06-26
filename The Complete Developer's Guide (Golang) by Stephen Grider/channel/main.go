package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// create a channel with type string.
	// Only values of type string can be sent into this channel.
	c := make(chan string)

	for _, link := range links {
		// wait for check link to finish.
		// checkLink(link)

		// run this function in a new go routine.
		// no need to wait for it to finish.
		// every time we call `go` keyword, it create a new go routine.
		// go checkLink(link)

		go checkLink(link, c)
	}

	// note: waiting for a value to be sent into the channel is a blocking call.
	// fmt.Println(<-c)

	// waiting to receive messages from all the go routines.
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// Watch `c` channel for any new value. As soon as we get a new value, pass it to `l`.
	for l := range c {
		// Pause the current go routine for 5 seconds.
		// time.Sleep(5 * time.Second)
		// go checkLink(l, c)

		// Function literal (anonymous function).
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) // Execute the function literal immediately.
	}
}

// func checkLink(link string) {
// 	_, err := http.Get(link)

// 	if err != nil {
// 		fmt.Println(link, "might be down!")
// 		return
// 	}

// 	fmt.Println(link, "is up!")
// }

// define a channel as a parameter to communicate between go routines.
func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		// Send a message into the channel.
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	// Send a message into the channel.
	c <- link
}
