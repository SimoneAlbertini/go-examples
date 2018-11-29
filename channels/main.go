package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	sites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	channel := make(chan string)

	for _, site := range sites {
		go checkLink(site, channel)
	}

	// for i := 0; i < len(sites); i++ {
	// 	fmt.Println(<-channel)
	// }

	for link := range channel {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, channel)
		}(link)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is probably down!")
		c <- link
		return
	}

	fmt.Println(link, "is up and running!")
	c <- link
}
