package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/buger/goterm"
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
		time.Sleep(2 * time.Second)
		go checkUrl(websites[i%len(websites)], channel)
		fmt.Print("\r", <-channel, " - ", i+1, " tries")
		goterm.Flush()
	}

}

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
