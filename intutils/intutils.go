package intutils

import "fmt"

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func intutils() {
	a := IntMin(2, 3)
	fmt.Println(a)
}

// func main() {
// 	fmt.Println("in main")
// }
