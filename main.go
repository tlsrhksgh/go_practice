package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("request failed")

func main() {
	var results = make(map[string]string)
	urls := []string{
		"https://www.airbnb.com",
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://academy.nomadcoders.co",
	}

	for _, url := range urls {
		result := "OK"
		err := hitUrl(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitUrl(url string) error {
	fmt.Println("Checking: ", url)
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		return errRequestFailed
	}

	return nil
}
