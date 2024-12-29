package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func getStatusCode(endpoint string) {
	defer wg.Done()
	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("Status code: %d and endpoint %s\n", resp.StatusCode, endpoint)
}

func main() {
	endpoints := []string{"http://www.google.com", "http://www.facebook.com", "http://www.twitter.com"}
	for _, endpoint := range endpoints {
		wg.Add(1)
		go getStatusCode(endpoint)
	}
	wg.Wait()
}