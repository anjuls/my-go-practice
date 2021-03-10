package main

import "fmt"

func main(){
	c := make(chan int, 2) // 2 is the buffer size
	c <- 1
	fmt.Println(<-c)
	c <- 2
	c <- 3
	fmt.Println(<-c)
	fmt.Println(<-c)
}