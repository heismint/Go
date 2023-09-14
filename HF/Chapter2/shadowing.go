package main

import (
	"fmt"
)

func main() {
	count := 12
	suffix := "minutes of bonus footage"
	format := "DVD"
	languages := append([]string{}, "Espanol")
	fmt.Println(count, suffix, "on", format, languages)
}
