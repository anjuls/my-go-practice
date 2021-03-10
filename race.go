package main

import (
	"fmt"
	"sync"
)

func main() {

	var data int
	var memoryAccess sync.Mutex
	go func() {
		memoryAccess.Lock()
		data++
		memoryAccess.Unlock()
	}()

	memoryAccess.Lock()
	if data == 0 {
		fmt.Println("the number is 0")
	} else {
		fmt.Printf("the value is %v\n.", data)
	}
	memoryAccess.Unlock()
}
