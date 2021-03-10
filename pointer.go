package main

import "fmt"

func main() {
	var p *int
	var i int
	i = 3
	p = &i
	fmt.Println("%v %d", p, (*p))
}
