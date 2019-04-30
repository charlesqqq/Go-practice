package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	c := make(chan string)
	for _, url := range urls {
		go checkURL(url, c)
	}
	for link := range c {
		go func(l string) {
			time.Sleep(3 * time.Second)
			checkURL(l, c)
		}(link)
	}
}

func checkURL(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "error")
		c <- url
		return
	}
	fmt.Println(url, "is ok")
	c <- url
}
