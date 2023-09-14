package main

import "fmt"

/*func main() {
	var greeting string = "Odogwu you bad"
	var oriki string = "baaje"
	fmt.Println(greeting,"\n" ,oriki)
}

var trying int = 32
var musa string = "Hello World"
var lie bool = false
var dec float64 = 77.90*/

// Declaring variables in a short form
/*func main() {
	var name string = "Golang"
	score := 67
	score = 88
	fmt.Printf("Hey, %v! You have scored %d/100 in Physics", name, score)
}*/

// Variable Scope
func main() {
	state := "Canada"
	{
		country := "USA"
		fmt.Println(country)
		fmt.Println(state)
	}
	fmt.Println(state)
}
