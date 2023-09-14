package main

import (
	"fmt"
	//"reflect" 
)

func main() {
	
	quantity := 4
	length, width := 1.34, 1.67
	customerName := "Nimota"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")


	/*fmt.Println("Cannonball!!!!")
    fmt.Println(reflect.TypeOf(25))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf(5.2))
	fmt.Println(reflect.TypeOf(1))
	fmt.Println(reflect.TypeOf(false))
	fmt.Println(reflect.TypeOf(1.0))
	fmt.Println(reflect.TypeOf("Hello"))*/
}
