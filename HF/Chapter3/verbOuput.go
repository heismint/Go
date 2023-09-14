// Verbs and Output in Go
package main

import (
	"fmt"
)

func main() {
	fmt.Printf("This %s cost %d cents each.\n", "gumballs", 23)
	fmt.Printf("That will be $%f please.\n", 0.23*5)
	//fmt.Println("I gat you bro")
	fmt.Printf("A float: %f\n", 3.1415)
	fmt.Print("An integer: %d\n", 20)
	fmt.Printf("A boolean: %t\n", true)
	fmt.Printf("A string: %s\n", "Omade")
	fmt.Printf("Values: %v %v %v\n", 1.2, "\t", true)
	fmt.Printf("Values: %#v %#v %#v\n", 1.2, "\t", true)
	fmt.Printf("Types: %T %T %T\n", 1.2, "\t", true)
	fmt.Printf("Percent sign: %%\n")

	fmt.Printf("%12s  |  %s\n", "Product", "Cost in cents")

	fmt.Printf("%12s  |  %2d\n", "Stamps", 50)
	fmt.Printf("%12s  |  %2d\n", "Paper clips", 5)
	fmt.Printf("%12s  |  %2d\n", "Tape", 99)
}
