package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://www.instagram.com",
		"https://www.twitter.com",
		"https://www.fb.com",
	}

	for _, link := range links {
		statusCheck(link)
	}

}

func statusCheck(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
		return
	}
	fmt.Println(link, "is up")
}
