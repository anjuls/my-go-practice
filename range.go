package main

import "fmt"

var pow_s = []int{1,2,4,8,16,32,64,128}

func main(){
	for i, v := range pow_s {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow := make([]int, 10)
    for i := range pow {
        pow[i] = 1 << uint(i)
    }
    
    for _, value := range pow {
        fmt.Printf("%d\n", value)
	}


cities := map[string]int{
		"New York":    8336697,
		"Los Angeles": 3857799,
		"Chicago":     2714856,
	}
	for key, value := range cities {
		fmt.Printf("%s has %d inhabitants\n", key, value)
	}

}