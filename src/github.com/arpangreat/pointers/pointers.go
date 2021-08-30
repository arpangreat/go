package main

import (
	"net/http"
)

func main() {
	client := &http.Client{}
	resp, err := client.Get("http://golangbootcamp.com")
}
