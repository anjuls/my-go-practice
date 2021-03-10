package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Unix()
	mins := now % 2
	fmt.Println(now)
	switch mins {
	case 0:
		fmt.Println("even")
	case 1:
		fmt.Println("odd")
	}
}