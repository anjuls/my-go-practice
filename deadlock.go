package main

import (
	"fmt"
	"sync"
	"time"
)

type Value struct {
	mu    sync.Mutex
	value int
}

func main() {

	var wg sync.WaitGroup
	var a, b Value

	printSum := func(v1, v2 *Value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()

}
