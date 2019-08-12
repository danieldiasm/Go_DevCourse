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

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//Added on lesson 78
	for l := range c {
		go func(link string) {
			//Added on lesson 79
			time.Sleep(time.Second * 3)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, error := http.Get(link)
	if error != nil {
		fmt.Println(link, " might be down!")
		c <- link
		return
	}

	fmt.Println("Link ", link, " seems to be okay!")
	c <- link
}
