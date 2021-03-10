package main

import "fmt"
import "math"

//constants
const Pi = 3.1428


// function and switch example
func location(city string) (string, string) {
	var region, continent string
	switch city {
	case "Indore", "Bhopal", "Ratlam":
		region, continent = "India", "Asia"
	case "Melbourne":
		region, continent = "Victoria", "Australia"
	default:
		region, continent = "unknown", "unknown"
	}
	return region, continent
}

func pointer_example(i *int)(*int){
	*i++
	return i
}


func main() {

	fmt.Println("Hello World")
	
	// implicit variables

	name, city := "Anjul", "Indore"
	age := 35
	fmt.Printf("%s (%d) of %s\n", name, age, city)

	// variable can have functions

	action := func(){
		fmt.Println("I am inside a function, Pi=", Pi)
	}

	action()

	fmt.Println("Value of Pi from Math Lib is ", math.Pi)

	region, continent := location(city)

	fmt.Printf("%s lives in %s, %s\n", name, region, continent)

	// pointer example
	age2 := pointer_example(&age) // 36 
	age2 = pointer_example(&age) //37
	age2 = pointer_example(&age) //38
	fmt.Printf("Memory:\n%d \n", *age2)
	


}
