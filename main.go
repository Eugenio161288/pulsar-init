package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Transport: &http.Transport{
		DialContext: (&net.Dialer{
			// Connect timeout
			Timeout: 5 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 5 * time.Second,
	},
	// Max total timeout
	Timeout: 60 * time.Second,
}

func main() {
	fmt.Println("Test google.com...")
	resp, err := httpClient.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Response status code: %v\n", resp.StatusCode)
	}
}
