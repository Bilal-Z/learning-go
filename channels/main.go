package main

import (
	"fmt"
	"net/http"
	"time"
)

func main()  {
	links := []string{
		"http://google.com",
		"http://reddit.com",
		"http://stackoverflow.com",
		"http://go.dev",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c) // create new go routine
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			go checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string){
	_, err := http.Get(link) // blocking call
	if err != nil{
		fmt.Println(link, "might be down")
		c <- link
		return
	}

	fmt.Println(link, "is up")
	c <- link
}