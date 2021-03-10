package main

import "fmt"
import "strings"

func WordCount(s string) map[string]int {
    words := strings.Fields(s)
    count := map[string]int{}
    fmt.Printf("Words: %#v \n",words)
    fmt.Printf("Count: %#v \n",count)
    for _, word := range words{
    	count[word]++
    	fmt.Printf("Count: %#v \n",count)
    }

    return count

}

func main() {
	m := WordCount("This is so amazing, I love it so much")
	fmt.Println(m)

}