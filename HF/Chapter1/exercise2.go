package main

import "fmt"

// A Go code that calculates total price with tax and determine if we have enough funds to make a purchase.

func main() {
	price := 100
	fmt.Println("Price is", price, "dollars")

	taxRate := 0.08
	tax := float64(price) * taxRate
	fmt.Println("Tax is", tax, "dollars.")

	total := float64(price) + tax
	fmt.Println("Total cost is", total, "dollars.")

	availableFunds := float64(120)
	fmt.Println(availableFunds, "dollars available.")
	fmt.Println("Within budget?", total <= availableFunds)
}
