package main

import (
	"fmt"
	//"strconv"
)

var a = "a"

func main() {
	a = "a"
	b := "b"
	if true {
		c := "c"
		if true {
			d := "d"
			fmt.Println(a)
			fmt.Println(b)
			fmt.Println(c)
			fmt.Println(d)
		}
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
	}
	fmt.Println(a)
	fmt.Println(b)

	/*z, err := strconv.ParseFloat("1.45", 64)
	y, err := strconv.ParseFloat("4.67", 64)
	fmt.Println(z, y, err)*/
}
