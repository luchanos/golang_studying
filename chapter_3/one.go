package main

import (
	"fmt"
	"log"
)

func paintNeeded(height float64, width float64) (result float64, err error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	area := width * height
	return area / 10, nil
}

func main() {
	amount, err := paintNeeded(4.2, -3.0)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%.2f liters needed\n", amount)
	}
}
