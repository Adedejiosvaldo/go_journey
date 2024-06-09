package main

import (
	"fmt"
	"net/http"
)

// type websites []string

func main() {

	links := []string{
		"http://www.google.com",
		"http://www.amazon.com",
		"http://www.facebook.com",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "Might be down")
		return
	}

	fmt.Println(link, "Is Up")
}
