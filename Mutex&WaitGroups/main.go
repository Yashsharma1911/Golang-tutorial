package main

import (
	"fmt"
	"net/http"
)

func main() {
	var site = "google.com"
	res, err := http.Get(site)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	}
	fmt.Printf("%d status code for %s", res.StatusCode, site)
}
