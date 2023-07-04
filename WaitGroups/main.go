package main

import (
	"fmt"
	"net/http"
	"sync"
)

//Basically we defines Wait group as per shown bellow

var wg sync.WaitGroup //this should be pointer but we are not using it as pointer now.

func main() {

	sites := []string{
		"https://github.com",
		"https://google.com",
		"https://linkedin.com",
		"https://instagram.com",
	}

	for _, site := range sites {
		go statusCode(site)
		wg.Add(1) // Add() method takes number of Goroutines that can be added in waitgroup
	}

	wg.Wait() // Wait() function dont allow to complete execution till all Goroutines comes back.

}

func statusCode(site string) {

	defer wg.Done() //Done() method notifies that execution of goroutine is over. also used 'defer' this line will execute at the end of this function.

	//As we know from previous tutorial of Goroutines that we have to tell to main function to wait for goroutine till it comes back
	//for that matter we are using sleep() method but WaitGroups give us freedom from using sleep() method

	res, err := http.Get(site)

	if err != nil {
		panic("Error In website :(")
	}
	fmt.Printf("%d is status code for %s\n", res.StatusCode, site)
}
