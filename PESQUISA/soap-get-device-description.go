package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	deviceDescURL := "http://192.168.1.1:1900/device.xml" // Replace with actual URL

	resp, err := http.Get(deviceDescURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body)) // This XML contains service URLs
}
