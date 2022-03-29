package main

import "fmt"

func paintNeeded(height float64, width float64) float64 {
	area := width * height
	return area / 10
}

func main() {
	var total, amount float64
	amount = paintNeeded(4.2, 3.0)
	total += amount
	amount = paintNeeded(5.2, 3.5)
	total += amount
	fmt.Printf("%.2f liters needed\n", total)
}
