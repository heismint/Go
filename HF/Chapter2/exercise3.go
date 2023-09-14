package main

import (
	"fmt"
)

func main() {
	for x := 1; x <= 3; x++ {
		fmt.Println(x)
	}
	for y := 1; y <= 3; y++ {
		fmt.Println(y)
	}
	for z := 2; z <= 3; z++ {
		fmt.Println(z)
	}
	for a := 1; a <= 3; a+= 2 {
		fmt.Println(a)
	}
	for b := 3; b >= 1; b-- {
		fmt.Println(b)
	}
	for c := 1; c < 3; c++ {
		fmt.Println(c)
	}
	for d := 1; d >= 3; d++ {
		fmt.Println(d)
	}
}
