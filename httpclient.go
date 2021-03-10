package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	resp, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status: ", resp.Status)
	fmt.Println("Response Header: ", resp.Header)
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < int(resp.ContentLength); i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
