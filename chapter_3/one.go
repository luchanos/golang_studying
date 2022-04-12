package main

import (
	"errors"
	"fmt"
	"log"
)

func paintNeeded(height float64, width float64) float64 {
	if height < 0 || width < 0 {
		err := errors.New("Value can't be negative!")
		log.Fatal(err)
	}
	area := width * height
	return area / 10
}

func main() {
	var total, amount float64
	amount = paintNeeded(4.2, -3.0)
	total += amount
	fmt.Printf("%.2f liters needed\n", total)
}
