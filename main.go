package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	websites := []string{
		// "http://google.com",
		// "http://amazon.com",
		// "http://stackoverflow.com",
		// "http://golang.org",
		// "https://guide.michelin.com/en/restaurants",
	}

	channel := make(chan string)

	/* Add argument websites to the main list */
	for i := 1; i < len(os.Args); i++ {
		url := os.Args[i]
		websites = append(websites, url)
	}

	/* Main loop */
	for i := 0; true; i++ {
		/* Sleep so it is not too suspicious network
		traffic from the website's viewpoint */
		time.Sleep(2 * time.Second)

		/* Launch the thread to do the checkup */
		go checkUrl(websites[i%len(websites)], channel)

		/* Whenever a thread finishes the request and
		responds into the channel, print their returned
		value. */
		fmt.Print(<-channel)
	}

}

/**
@param channel - Channels in go are like a pipe in C.
	All threads will, this way, send their response in
	the same channel and the main thread will then be
	waiting for responses independant of who sent.

@param url - Website's url intended to check.
*/
func checkUrl(url string, channel chan string) {
	if isWebsiteUp(url) {
		channel <- (url + " is up!")
	} else {
		channel <- (url + " is down...")
	}
}

func isWebsiteUp(url string) bool {
	response, err := http.Get(url)

	if err != nil {
		return false
	}

	return response.StatusCode == 200
}
